package expects

type IExpect interface {
	error
	GetMessage() string
	GetCode() int
}
type StandardExpect struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewStandardExpect(message string, code int) *StandardExpect {
	return &StandardExpect{
		Message: message,
		Code:    code,
	}
}

func (e *StandardExpect) Error() string {
	return e.Message
}

func (e *StandardExpect) GetCode() int { return e.Code }

func (e *StandardExpect) GetMessage() string {
	return e.Message
}
