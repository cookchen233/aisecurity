package handlers

import (
	"aisecurity/types"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"log"
)

type Handler struct {
	HandlerService types.IService
}

func Convert(handler types.IHandler, handerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.SetContext(ctx)
		handerFunc(ctx)
	}
}

func (handler *Handler) SetContext(ctx context.Context) {
	handler.HandlerService.SetContext(ctx)
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
		//return errors.New(fmt.Sprint(errMsgs))
		return err
	}
	return nil
}
