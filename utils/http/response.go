package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Payload struct {
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
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
			Value interface{} `json:"Value"`
		}{data}
	case reflect.Array, reflect.Slice:
		respData = struct {
			List interface{} `json:"List"`
		}{data}
	default:
		varValue := reflect.ValueOf(data)
		if varValue.Kind() == reflect.Struct {
			if _, ok := varValue.Type().FieldByName("Item"); !ok {
				respData = struct {
					Item interface{} `json:"Item"`
				}{data}
			} else {
				respData = data
			}
		}
	}
	c.JSON(http.StatusOK, Payload{
		http.StatusOK,
		"成功",
		respData,
	})
}

func SuccessWithCode(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, Payload{
		code,
		"成功",
		data,
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
	c.AbortWithStatusJSON(statusCode, Payload{
		code,
		err.Error(),
		nil,
	})
}
