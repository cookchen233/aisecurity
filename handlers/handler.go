package handlers

import (
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/utils/http"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"log"
)

type IHandler interface {
	SetContext(ctx context.Context)
	Create(c *gin.Context)
	GetList(c *gin.Context)
	GetDetails(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	Service services.IService
	Filter  structs.IFilter
	Entity  structs.IEntity
}

func Convert(handler IHandler, handerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.SetContext(c)
		handerFunc(c)
	}
}

func (handler *Handler) SetContext(ctx context.Context) {
	handler.Service.SetContext(ctx)
}

func (handler *Handler) Validate(data interface{}) error {
	locale := zh_Hans_CN.New()
	uni := ut.New(locale, locale)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	valid := validator.New()
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

func (handler *Handler) Create(c *gin.Context) {
	if err := c.ShouldBindJSON(handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.Service.Create(handler.Entity)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *Handler) Update(c *gin.Context) {
	if err := c.ShouldBindJSON(handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.Service.Update(handler.Entity)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *Handler) GetList(c *gin.Context) {
	if err := c.ShouldBindQuery(handler.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	list, err := handler.Service.GetList(handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

func (handler *Handler) GetDetails(c *gin.Context) {
	if err := c.ShouldBindQuery(handler.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	details, err := handler.Service.GetDetails(handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, details)
}

func (handler *Handler) Delete(c *gin.Context) {
	if err := c.ShouldBindJSON(handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	err := handler.Service.Delete(handler.Entity)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, nil)
}
