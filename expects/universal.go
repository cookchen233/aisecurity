package expects

type NotImplementedMethod struct {
	StandardExpect
	Message string
}

func NewNotImplementedMethod() *NotImplementedMethod {
	e := &NotImplementedMethod{
		Message: "该功能暂未上线",
	}
	e.StandardExpect.Message = e.Message
	return e
}

type DataNotFound struct {
	StandardExpect
	Message string
}

func NewDataNotFound() *DataNotFound {
	e := &DataNotFound{
		Message: "该记录未找到或已被删除",
	}
	e.StandardExpect.Message = e.Message
	return e
}

type FileUploadError struct {
	StandardExpect
	Message string
	Size    int64
}

func NewFileUploadError(message string) *FileUploadError {
	e := &FileUploadError{
		Message: message,
	}
	e.StandardExpect.Message = e.Message
	return e
}

type EmptyData struct {
	StandardExpect
	Message string
}

func NewEmptyData() *DataNotFound {
	e := &DataNotFound{
		Message: "数据为空",
	}
	e.StandardExpect.Message = e.Message
	return e
}
