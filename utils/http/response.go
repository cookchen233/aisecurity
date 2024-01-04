package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strings"
)

type Payload struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Ref     string      `json:"ref"`
	Traceid string      `json:"traceid"`
}

func Success(c *gin.Context, data interface{}) {
	varType := reflect.TypeOf(data).Kind()
	// Checking if the variable is a scalar type
	var respData interface{}
	switch varType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Bool,
		reflect.String:
		respData = struct {
			Value interface{} `json:"value"`
		}{data}
	case reflect.Array, reflect.Slice:
		respData = struct {
			List interface{} `json:"list"`
		}{data}
	default:
		varValue := reflect.ValueOf(data)
		if varValue.Kind() == reflect.Struct {
			if _, ok := varValue.Type().FieldByName("Item"); !ok {
				respData = struct {
					Item interface{} `json:"item"`
				}{data}
			} else {
				respData = data
			}
		} else {
			respData = data
		}
	}
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	c.JSON(http.StatusOK, Payload{
		http.StatusOK,
		"成功",
		respData,
		"",
		traceid.(string),
	})
}

func SuccessWithCode(c *gin.Context, data interface{}, code int) {
	traceid, _ := c.Get("traceid")
	if traceid == nil {
		traceid = ""
	}
	c.JSON(http.StatusOK, Payload{
		code,
		"成功",
		data,
		"",
		traceid.(string),
	})
}

func Error(c *gin.Context, err error, code int) {
	var statusCode = c.Writer.Status()
	var errMessage string
	if !c.Writer.Written() {
		statusCode = 400
		errMessage = "请求数据错误, 请检查"
		if code >= 1000 {
			statusCode = 500
			errMessage = "服务器发生错误, 请稍后再试"
		}
		c.Status(statusCode)
	}
	c.Error(errors.WithStack(err))
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	if err, ok := err.(interface{ Cause() error }); ok {
		errMessage = err.Cause().Error()
	}

	cause := errors.Cause(err)
	// Convert the full error message and the cause into strings
	fullErrMsg := err.Error()
	causeErrMsg := cause.Error()
	if fullErrMsg != causeErrMsg {
		// Remove the cause part from the full error message
		errMessage = strings.TrimSuffix(fullErrMsg, causeErrMsg)
		// Trim any leading colon or space
		errMessage = strings.TrimRight(errMessage, ": ")
	}

	c.AbortWithStatusJSON(statusCode, Payload{
		code,
		errMessage,
		nil,
		fullErrMsg,
		traceid.(string),
	})
}
