package expects

type IExpect interface {
	error
	ExpectedError() string
}
type StandardExpect struct {
	Message string `json:"message"`
}

func (e *StandardExpect) Error() string {
	return e.Message
}

func (e *StandardExpect) ExpectedError() string {
	return e.Error()
}
