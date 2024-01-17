package http

import (
	"aisecurity/expects"
	"aisecurity/properties"
	"aisecurity/utils"
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
	data = prepareData(data)
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	c.JSON(http.StatusOK, Payload{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    data,
		Traceid: traceid.(string),
	})
	if len(c.Errors) > 0 {
		Error(c, c.Errors[0], 1000)
	}
}

func prepareData(data interface{}) interface{} {
	varValue := reflect.ValueOf(data)
	varType := reflect.TypeOf(data)

	if varType == nil {
		return struct {
			List []interface{} `json:"list"`
		}{[]interface{}{}}
	}

	switch varType.Kind() {
	case reflect.Slice, reflect.Array:
		if varValue.Len() == 0 {
			return struct {
				List interface{} `json:"list"`
			}{[]interface{}{}}
		}
		return struct {
			List interface{} `json:"list"`
		}{data}
	case reflect.Struct:
		return prepareStructData(varValue, varType, data)
	case reflect.Pointer:
		return preparePointerData(varValue, data)
	default:
		return struct {
			Value interface{} `json:"value"`
		}{data}
	}
}

func prepareStructData(varValue reflect.Value, varType reflect.Type, data interface{}) interface{} {
	// Ensure varValue is addressable, necessary for setting field values
	if !varValue.CanAddr() {
		// Create a copy that is addressable
		copy := reflect.New(varType).Elem()
		copy.Set(varValue)
		varValue = copy
	}

	if listField := varValue.FieldByName("List"); listField.IsValid() && listField.Kind() == reflect.Slice {
		if listField.IsNil() || listField.Len() == 0 {
			// Set the List field to an empty slice
			listField.Set(reflect.MakeSlice(listField.Type(), 0, 0))
		}
		return varValue.Interface()
	}

	if _, ok := varType.FieldByName("Item"); ok {
		return data
	}

	return struct {
		Item interface{} `json:"item"`
	}{data}
}

func preparePointerData(varValue reflect.Value, data interface{}) interface{} {
	if varValue.IsNil() {
		return struct {
			Item interface{} `json:"item"`
		}{nil}
	}
	dereferencedValue := varValue.Elem()
	if dereferencedValue.Kind() == reflect.Struct {
		return prepareStructData(dereferencedValue, dereferencedValue.Type(), dereferencedValue.Interface())
	}
	return dereferencedValue.Interface()
}

func SuccessWithCode(c *gin.Context, data interface{}, code properties.ResponseCode) {
	data = prepareData(data)
	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	c.JSON(http.StatusOK, Payload{
		Code:    code,
		Message: "成功",
		Data:    data,
		Traceid: traceid.(string),
	})
	if len(c.Errors) > 0 {
		Error(c, c.Errors[0], 1000)
	}
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

	// add error to context
	if len(c.Errors) == 0 || !errors.Is(err, c.Errors[len(c.Errors)-1]) {
		c.Error(utils.ErrorWithStack(err))
		//c.Error(err)
	}

	// if specified code
	var errMessage = code.Message()

	// if wrapped messages
	cause := errors.Cause(err)
	fullErrMsg := err.Error()
	causeErrMsg := cause.Error()
	if fullErrMsg != causeErrMsg {
		errMessage = strings.TrimRight(strings.TrimSuffix(fullErrMsg, causeErrMsg), ": ")
	}

	// if err is IExpect type
	if code == properties.ServerError || code == properties.RequestError {
		var expect expects.IExpect
		if errors.As(err, &expect) {
			errMessage = expect.ExpectedError()
		}
	}

	traceid, ex := c.Get("traceid")
	if !ex {
		traceid = ""
	}
	c.AbortWithStatusJSON(statusCode, Payload{
		code,
		errMessage,
		nil,
		fullErrMsg,
		traceid.(string),
	})
}
