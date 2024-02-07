package expects

import "github.com/pkg/errors"

type ServerError struct {
	StandardExpect
}

func NewServerError(message string) *NotImplementedMethod {
	e := &NotImplementedMethod{
		StandardExpect{
			Message: message,
			Code:    2000,
		},
	}
	return e
}

type NotImplementedMethod struct {
	StandardExpect
}

func NewNotImplementedMethod() *NotImplementedMethod {
	e := &NotImplementedMethod{
		StandardExpect{
			Message: "该功能暂未上线",
			Code:    2001,
		},
	}
	return e
}

type DataNotFound struct {
	StandardExpect
}

func NewDataNotFound() *DataNotFound {
	e := &DataNotFound{
		StandardExpect{
			Message: "该记录未找到或已被删除",
			Code:    2002,
		},
	}
	return e
}
func IsDataNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *DataNotFound
	return errors.As(err, &e)
}

type EmptyData struct {
	StandardExpect
}

func NewEmptyData() *DataNotFound {
	e := &DataNotFound{
		StandardExpect{
			Message: "数据为空",
			Code:    2003,
		}}
	return e
}

type InValidID struct {
	StandardExpect
}

func NewInValidID() *InValidID {
	e := &InValidID{
		StandardExpect{
			Message: "无效的ID参数",
			Code:    2004,
		},
	}
	return e
}
