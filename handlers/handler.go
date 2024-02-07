package handlers

import (
	"aisecurity/expects"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	http2 "net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type IHandler interface {
	SetRequestContext(c *gin.Context, handler IHandler)
	GetService(c *gin.Context) services.IService
	GetFilter(c *gin.Context) structs.IFilter
	GetEntity(c *gin.Context) structs.IEntity
	GetValidator() *validator.Validate
}

type Handler struct {
	Service   services.IService
	Filter    structs.IFilter
	Entity    structs.IEntity
	Validator *validator.Validate
}

func Convert(handler IHandler, handerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.SetRequestContext(c, handler)
		handerFunc(c)
	}
}

func (h *Handler) GetService(c *gin.Context) services.IService {
	utils.Logger.Error("called empty method GetService", zap.String("url", c.Request.RequestURI))
	return h.Service
}
func (h *Handler) GetFilter(c *gin.Context) structs.IFilter {
	utils.Logger.Error("called empty method GetFilter", zap.String("url", c.Request.RequestURI))

	return h.Filter
}
func (h *Handler) GetEntity(c *gin.Context) structs.IEntity {
	utils.Logger.Error("called empty method GetEntity", zap.String("url", c.Request.RequestURI))
	return h.Entity
}
func (h *Handler) GetValidator() *validator.Validate {
	return nil
}

func (h *Handler) SetRequestContext(c *gin.Context, childHandler IHandler) {
	h.Service = childHandler.GetService(c)
	h.Filter = childHandler.GetFilter(c)
	h.Entity = childHandler.GetEntity(c)
	h.Validator = childHandler.GetValidator()
}

func (h *Handler) Validate(data interface{}) error {
	locale := zh_Hans_CN.New()
	uni := ut.New(locale, locale)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	valid := h.Validator
	if valid == nil {
		valid = validator.New()
	}
	if err := zh.RegisterDefaultTranslations(valid, trans); err != nil {
		log.Printf("RegisterDefaultTranslations error: %v", err)
	}

	if err := valid.Struct(data); err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, e.Translate(trans))
		}
		return fmt.Errorf("%v", errMsgs)
	}
	return nil
}

func (h *Handler) UploadFile(c *gin.Context, basePath string, fileHeaders []*multipart.FileHeader, maxSize int64, allowedMimeTypes *types.AllowedMimeTypes) ([]types.UploadedFile, error) {
	// Assuming the save_dir parameter is sent in a POST request
	saveDir := c.PostForm("save_dir")
	// Resolve the intended path
	intendedPath := filepath.Join(basePath, filepath.Clean(saveDir))
	// Ensure the resolved path is within the desired base directory
	basePath, err := filepath.Abs(basePath)
	if err != nil {
		return nil, errors.WithStack(expects.NewFileUploadDenied("指定的上传目录有误, " + basePath))
	}
	absPath, err := filepath.Abs(intendedPath)
	if err != nil {
		return nil, errors.WithStack(expects.NewFileUploadDenied("指定的上传目录有误2, " + absPath))
	}
	rel, err := filepath.Rel(basePath, absPath)
	if err != nil {
		return nil, errors.WithStack(expects.NewFileUploadDenied("指定的上传目录有误3, " + rel))
	}
	if strings.Contains(rel, "..") {
		return nil, errors.WithStack(expects.NewFileUploadDenied("指定的上传目录有误4, " + rel))
	}

	var uploadedFiles []types.UploadedFile
	for _, fh := range fileHeaders {
		// Check the file size and save the file as usual
		if fh.Size > maxSize {
			return nil, errors.WithStack(expects.NewFileUploadDenied("文件大小超过了 " + gfile.FormatSize(maxSize)))
		}
		// Detect the file type
		buffer, err := utils.ReadFileBuffer(fh)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if !allowedMimeTypes.IsAllowed(http2.DetectContentType(buffer)) {
			return nil, errors.WithStack(expects.NewFileUploadDenied("仅支持: " + strings.Join(allowedMimeTypes.Types, ", ") + " 文件格式"))
		}

		// Create a new seeded source
		src := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(src)
		// Get the current date in Y-m and Ymd format
		now := time.Now()
		dirFormat := now.Format("2006-01")
		fileDateFormat := now.Format("200601021504")
		// Generate a 6-digit random number
		randomNumber := rnd.Intn(1000000)
		// Create the directory name and check if it exists
		dirName := filepath.Join(intendedPath, dirFormat)
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			// Create the directory if it does not exist
			err := os.MkdirAll(dirName, 0755) // 0755 permissions
			if err != nil {
				utils.Logger.Error("Error creating directory", zap.Error(err))
			}
		}
		// Form the file name
		name := fmt.Sprintf("%s-%06d%s", fileDateFormat, randomNumber, filepath.Ext(fh.Filename))
		filename := filepath.Join(dirName, name)

		if err := c.SaveUploadedFile(fh, filename); err != nil {
			return nil, errors.WithStack(err)

		}
		uploadedFiles = append(uploadedFiles, types.UploadedFile{Name: fh.Filename, URL: filename, Size: fh.Size, CreateTime: time.Now()})
	}
	return uploadedFiles, nil
}

func (h *Handler) GetRequestBody(c *gin.Context) string {
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	return string(bodyBytes)
}
