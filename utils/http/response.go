package http

import (
	"aisecurity/properties"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strings"
)

type Payload struct {
	Code    properties.ResponseCode `json:"code"`
	Message string                  `json:"message"`
	Data    interface{}             `json:"data"`
	Ref     string                  `json:"ref"`
	Traceid string                  `json:"traceid"`
}

func Success(c *gin.Context, data interface{}) {
	varType := reflect.TypeOf(data).Kind()
	// Checking if the variable is a scalar type
	switch varType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Bool,
		reflect.String:
		data = struct {
			Value interface{} `json:"value"`
		}{data}
	case reflect.Array, reflect.Slice:
		data = struct {
			List interface{} `json:"list"`
		}{data}
	default:
		varValue := reflect.ValueOf(data)
		if varValue.Kind() == reflect.Struct {
			if _, ok := varValue.Type().FieldByName("Item"); !ok {
				data = struct {
					Item interface{} `json:"item"`
				}{data}
			}
		} else if varValue.Kind() == reflect.Pointer {
			// Dereference the pointer to get the value it points to.
			dereferencedValue := varValue.Elem()
			// Check the kind of the dereferenced value.
			if dereferencedValue.Kind() == reflect.Struct {
				// Now check if the struct has a field named "Item".
				if _, ok := dereferencedValue.Type().FieldByName("Item"); !ok {
					data = struct {
						Item interface{} `json:"item"`
					}{data}
				}
			} else {
				// Handle other kinds (like pointer to a basic type) appropriately.
				data = dereferencedValue.Interface()
			}
		}
	}
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	c.JSON(http.StatusOK, Payload{
		http.StatusOK,
		"成功",
		data,
		"",
		traceid.(string),
	})
	if len(c.Errors) > 0 {
		Error(c, c.Errors[0], 1000)
	}
}

func SuccessWithCode(c *gin.Context, data interface{}, code properties.ResponseCode) {
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

func Error(c *gin.Context, err error, code properties.ResponseCode) {
	var statusCode = c.Writer.Status()
	if !c.Writer.Written() {
		statusCode = http.StatusBadRequest
		if code >= 1000 {
			statusCode = http.StatusInternalServerError
		}
		c.Status(statusCode)
	}
	if len(c.Errors) == 0 || err != c.Errors[len(c.Errors)-1] {
		c.Error(errors.WithStack(err))
	}
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}

	var errMessage = properties.ResponseCode(code).Message()
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
