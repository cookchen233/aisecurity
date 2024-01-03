package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

type Payload struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
		traceid.(string),
	})
}

func Error(c *gin.Context, err error, code int) {
	var statusCode = c.Writer.Status()
	if !c.Writer.Written() {
		statusCode = 400
		if code >= 1000 {
			statusCode = 500
		}
		c.Status(statusCode)
	}
	c.Error(errors.WithStack(err))
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	c.AbortWithStatusJSON(statusCode, Payload{
		code,
		err.Error(),
		nil,
		traceid.(string),
	})
}
