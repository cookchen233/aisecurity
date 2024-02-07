package ipc

import (
	"aisecurity/ent/dao"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/posts"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"regexp"
	"time"
)

type TwoHandler struct {
	handlers.IPCHandler
	Service *services.EventService
}

func NewTwoHandler() *TwoHandler { return &TwoHandler{} }
func (h *TwoHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *TwoHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *TwoHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *TwoHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewEventService(c)
	h.Filter = &filters.Event{}
	h.Entity = &entities.Event{}
	h.IPCHandler.SetRequestContext(c, h)
}

func (h *TwoHandler) ReportEvent(c *gin.Context) {
	// Get the body must be done before the body is read
	bodyString := h.GetRequestBody(c)
	pattern := `"((ImageData|ImageDataLabeled))"\s?:\s?"([^"]{64}).*([^"]{64})"`
	re := regexp.MustCompile(pattern)
	bodyString = re.ReplaceAllString(bodyString, `"$1": "$3...$4"`)

	var p posts.TwoEvent
	if err := c.ShouldBindJSON(&p); err != nil {
		http.Error(c, err, 1000)
		return
	}

	// event type
	eventType := enums.EventType(0).GetEventTypeByTwoType(p.Result.Type)

	// event time
	timestampMicro := int64(p.TimeStamp)
	eventTime := time.Unix(0, timestampMicro*int64(time.Microsecond))

	// description
	description := p.Result.Description
	if description == "" {
		description = eventType.Label()
	}

	// device
	deviceService := services.NewDevice(c)
	device, err := deviceService.GetBySn(p.BoardID)
	if err != nil {
		if dao.IsNotFound(err) {
			device, err = deviceService.Create(&entities.Device{
				Device: dao.Device{
					Brand:      enums.DB1,
					Model:      enums.DM1,
					Name:       "图为T1206",
					Sn:         p.BoardID,
					DeviceType: enums.DT1,
				},
			})
			if err != nil {
				http.Error(c, utils.ErrorWithStack(err), 2000)
				return
			}
		} else {
			http.Error(c, utils.ErrorWithStack(err), 2000)
			return
		}
	}
	d := device.(*entities.Device)

	// device installation
	deviceInstallationService := services.NewDeviceInstallationService(c)
	_, err = deviceInstallationService.GetByDeviceID(d.ID)
	if err != nil {
		if dao.IsNotFound(err) {
			area, _ := services.NewAreaService(c).GetAnyOne()
			locationData, _ := json.Marshal(p.GPS)
			_, _ = deviceInstallationService.Create(&entities.DeviceInstallation{
				DeviceInstallation: dao.DeviceInstallation{
					DeviceID:     d.ID,
					AreaID:       area.(*entities.Area).ID,
					AliasName:    "图为T1206默认安装",
					Longitude:    p.GPS.Longitude,
					Latitude:     p.GPS.Latitude,
					LocationData: string(locationData),
					Location:     "默认位置",
					InstallTime:  time.Now(),
				},
			})
		}
	}

	ent := entities.Event{
		Event: dao.Event{
			DeviceID:    d.ID,
			DataID:      p.AlarmID,
			EventTime:   eventTime,
			EventType:   eventType,
			Description: description,
			RawData:     bodyString,
		},
	}

	// Save the image
	basePath := "data/two/images"
	image, err := utils.Base64ToImage(basePath, p.ImageData)
	if err != nil {
		utils.Logger.Error("failed to convert base64 to image", zap.Error(err))
		http.Error(c, err, 2000)
		return
	}
	imageInfo, err := os.Stat(image)
	if err != nil {
		utils.Logger.Error("failed getting file info", zap.Error(err))
		http.Error(c, err, 2000)
		return
	}
	ent.Images = append(ent.Images, &types.UploadedImage{UploadedFile: types.UploadedFile{
		Name:       p.LocalRawPath,
		URL:        image,
		Size:       imageInfo.Size(),
		CreateTime: time.Now(),
	}})

	// Save the labeled image
	if p.ImageDataLabeled != "" {
		labeledImage, err := utils.Base64ToImage(basePath, p.ImageDataLabeled)
		if err != nil {
			utils.Logger.Error("failed to convert base64 to image", zap.Error(err))
			http.Error(c, err, 2000)
			return
		}
		labeledImageInfo, err := os.Stat(labeledImage)
		if err != nil {
			utils.Logger.Error("failed getting file info", zap.Error(err))
			http.Error(c, err, 2000)
			return
		}
		ent.LabeledImages = append(ent.LabeledImages, &types.UploadedImage{UploadedFile: types.UploadedFile{
			Name:       p.LocalLabeledPath,
			URL:        labeledImage,
			Size:       labeledImageInfo.Size(),
			CreateTime: time.Now(),
		}})
	}

	// after the videos being uploaded, the event will be updated.
	videoService := services.NewVideoService(c)
	video, err := videoService.CreateOrUpdateByName(&entities.Video{
		Video: dao.Video{
			Name: p.VideoFile,
		},
	})
	if err != nil {
		utils.Logger.Error("failed to create or update video", zap.Error(err))
	}
	if video != nil {
		ent.VideoID = video.(*entities.Video).ID
	}

	saved, err := h.Service.Create(&ent)
	if err != nil {
		http.Error(c, err, 2000)
		return
	}

	// event log
	eventLogService := services.NewEventLogService(c)
	_, err = eventLogService.Create(&entities.EventLog{
		EventLog: dao.EventLog{
			DeviceID: d.ID,
			EventID:  saved.(*entities.Event).ID,
			LogType:  enums.ELT1,
		},
	})
	http.Success(c, saved)
}

func (h *TwoHandler) UploadVideos(c *gin.Context) {
	var basePath = "./data/two/videos"
	var maxSize int64 = 1024 * 1024 * 100
	var allowedMimeTypes = types.NewAllowedMimeTypes([]string{})
	form, _ := c.MultipartForm()
	if form == nil {
		http.Error(c, utils.ErrorWithStack(expects.NewClientError("空的表单数据")))
		return
	}
	fileHeaders := form.File["Video"]
	uploadedFiles, err := h.UploadFile(c, basePath, fileHeaders, maxSize, allowedMimeTypes)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err))
		return
	}
	if len(uploadedFiles) == 0 {
		http.Error(c, utils.ErrorWithStack(expects.NewEmptyData()), 1000)
		return
	}
	var saveds []structs.IEntity
	var failures []string
	videoService := services.NewVideoService(c)
	for _, file := range uploadedFiles {
		saved, err := videoService.CreateOrUpdateByName(&entities.Video{
			Video: dao.Video{
				Name: file.Name,
				Size: file.Size,
				URL:  file.URL,
			},
		})
		if err != nil {
			failures = append(failures, fmt.Sprintf("创建或更新视频记录时发生错误, file.Name: %s, %v", file.Name, err))
			continue
		}
		saveds = append(saveds, saved)

	}
	http.Success(c, map[string]any{"saveds": saveds, "failures": failures})
}

func (h *TwoHandler) ExtraMessages(c *gin.Context) {
	http.Error(c, utils.ErrorWithStack(expects.NewNotImplementedMethod()), 1000)
}
