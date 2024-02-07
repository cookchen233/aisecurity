package http

import (
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

type Payload struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Ref     string      `json:"ref"`
	Traceid string      `json:"traceid"`
}

func Success(c *gin.Context, data interface{}) {
	data2 := prepareData(data)
	traceid := c.GetString("traceid")
	payload := Payload{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    data2,
		Traceid: traceid,
	}
	c.Set("payload", payload)
	c.JSON(http.StatusOK, payload)
	if len(c.Errors) > 0 {
		Error(c, c.Errors[0], 1000)
	}
}

func prepareData(data interface{}) interface{} {
	_, ok := data.(structs.IEntity)
	if ok {
		return struct {
			Item interface{} `json:"item"`
		}{data}
	}

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
	//case reflect.Struct:
	//	return prepareStructData(varValue, varType, data)
	//case reflect.Pointer:
	//	return preparePointerData(varValue, data)
	default:
		return data
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

func Error(c *gin.Context, err error, codes ...int) {
	code := 3000
	if len(codes) > 0 {
		code = codes[0]
	}
	var statusCode = c.Writer.Status()
	if !c.Writer.Written() {
		statusCode = http.StatusBadRequest
		if code >= 2000 {
			statusCode = http.StatusInternalServerError
		}
		c.Status(statusCode)
	}

	// add error to context
	if len(c.Errors) == 0 || !errors.Is(err, c.Errors[len(c.Errors)-1]) {
		c.Error(utils.ErrorWithStack(err))
		//c.Error(err)
	}

	var responseError expects.IExpect
	var expect expects.IExpect
	if errors.As(err, &expect) {
		responseError = expect
	} else {
		if code >= 2000 {
			responseError = expects.NewServerError("服务器发生错误, 请稍后再试")
		} else {
			responseError = expects.NewClientError("请求数据错误, 请检查")
		}
	}

	traceid := c.GetString("traceid")
	payload := Payload{
		responseError.GetCode(),
		responseError.GetMessage(),
		nil,
		err.Error(),
		traceid,
	}
	c.AbortWithStatusJSON(statusCode, payload)
	c.Set("payload", payload)

}
