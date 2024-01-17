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
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"regexp"
	"time"
)

type TwoHandler struct {
	handlers.IPCHandler
	Service *services.IPCReportEventService
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
func (h *TwoHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewIPCReportEventService()
	h.Service.Ctx = c
	h.Filter = &filters.IPCReportEvent{}
	h.Entity = &entities.IPCReportEvent{}
	h.IPCHandler.SetRequestContext(c, h)
}

func (h *TwoHandler) ReportEvent(c *gin.Context) {
	// Get the body must be done before the body is read
	bodyString := h.GetRequestBody(c)
	pattern := `"((ImageData|ImageDataLabeled))"\s?:\s?"([^"]{64}).*([^"]{64})"`
	re := regexp.MustCompile(pattern)
	bodyString = re.ReplaceAllString(bodyString, `"$1": "$3...$4"`)

	var p posts.TwoIPCReportEvent
	if err := c.ShouldBindJSON(&p); err != nil {
		http.Error(c, err, 900)
		return
	}

	eventType := enums.IPCReportEventType(0).GetEventTypeByTwoType(p.Result.Type)

	timestampMicro := int64(p.TimeStamp)
	eventTime := time.Unix(0, timestampMicro*int64(time.Microsecond))

	description := p.Result.Description
	if description == "" {
		description = eventType.Label()
	}

	ent := entities.IPCReportEvent{
		IPCReportEvent: dao.IPCReportEvent{
			DeviceID:    p.BoardID,
			EventID:     p.AlarmID,
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
		http.Error(c, err, 1000)
		return
	}
	imageInfo, err := os.Stat(image)
	if err != nil {
		utils.Logger.Error("failed getting file info", zap.Error(err))
		http.Error(c, err, 1000)
		return
	}
	ent.Images = append(ent.Images, types.UploadedImage{UploadedFile: types.UploadedFile{
		Name:      p.LocalRawPath,
		URL:       image,
		Size:      imageInfo.Size(),
		CreatedAt: time.Now(),
	}})

	// Save the labeled image
	if p.ImageDataLabeled != "" {
		labeledImage, err := utils.Base64ToImage(basePath, p.ImageDataLabeled)
		if err != nil {
			utils.Logger.Error("failed to convert base64 to image", zap.Error(err))
			http.Error(c, err, 1000)
			return
		}
		labeledImageInfo, err := os.Stat(labeledImage)
		if err != nil {
			utils.Logger.Error("failed getting file info", zap.Error(err))
			http.Error(c, err, 1000)
			return
		}
		ent.LabeledImages = append(ent.LabeledImages, types.UploadedImage{UploadedFile: types.UploadedFile{
			Name:      p.LocalLabeledPath,
			URL:       labeledImage,
			Size:      labeledImageInfo.Size(),
			CreatedAt: time.Now(),
		}})
	}

	// after the videos being uploaded, the event will be updated.
	ent.Videos = append(ent.Videos, types.UploadedVideo{UploadedFile: types.UploadedFile{
		Name: p.VideoFile,
	}})

	saved, err := h.Service.Create(&ent)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (h *TwoHandler) UploadVideos(c *gin.Context) {
	var basePath = "./data/two/videos"
	var maxSize int64 = 1024 * 1024 * 100
	var allowedMimeTypes = types.NewAllowedMimeTypes([]string{})
	form, _ := c.MultipartForm()
	fileHeaders := form.File["Video"]
	uploadedFiles, err, code := h.UploadFile(c, basePath, fileHeaders, maxSize, allowedMimeTypes)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), code)
		return
	}
	if len(uploadedFiles) == 0 {
		http.Error(c, utils.ErrorWithStack(expects.NewEmptyData()), 900)
		return
	}
	var saveds []structs.IEntity
	for _, file := range uploadedFiles {
		event, err := h.Service.GetByVideoName(file.Name)
		if err != nil {
			if dao.IsNotFound(err) {
				http.Error(c, utils.ErrorWithStack(fmt.Errorf("根据视频文件名未找到事件记录")), 900)
				return
			}
			http.Error(c, utils.ErrorWithStack(err), 1000)
			return
		}
		for _, v := range event.Videos {
			if v.Name == file.Name {
				v.UploadedFile = file
				break
			}
		}
		saved, err := h.Service.UpdateVideos(event)
		if err != nil {
			http.Error(c, err, 1000)
			return
		}
		saveds = append(saveds, saved)
	}
	http.Success(c, saveds)
}

func (h *TwoHandler) ExtraMessages(c *gin.Context) {
	http.Error(c, utils.ErrorWithStack(expects.NewNotImplementedMethod()), 900)
}
