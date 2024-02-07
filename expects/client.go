package expects

type ClientError struct {
	StandardExpect
}

func NewClientError(message string) *DataNotFound {
	e := &DataNotFound{
		StandardExpect{
			Message: message,
			Code:    1000,
		}}
	return e
}

type NotLoggedIn struct {
	StandardExpect
}

func NewNotLoggedIn() *DataNotFound {
	e := &DataNotFound{
		StandardExpect{
			Message: "未登录",
			Code:    1001,
		}}
	return e
}

type Unauthorized struct {
	StandardExpect
}

func NewUnauthorized() *DataNotFound {
	e := &DataNotFound{
		StandardExpect{
			Message: "未授权",
			Code:    1002,
		}}
	return e
}

type FileUploadDenied struct {
	StandardExpect
}

func NewFileUploadDenied(message string) *FileUploadDenied {
	e := &FileUploadDenied{
		StandardExpect{
			Message: message,
			Code:    1003,
		},
	}
	return e
}

type AlreadyExistedUsername struct {
	StandardExpect
}

func NewAlreadyExistedUsername() *FileUploadDenied {
	e := &FileUploadDenied{
		StandardExpect{
			Message: "已存在的登录名",
			Code:    1004,
		},
	}
	return e
}
