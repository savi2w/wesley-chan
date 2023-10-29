package errors

type JsonError struct {
	Message string `json:"message"`
}

func Wrap(err error) *JsonError {
	return &JsonError{
		Message: err.Error(),
	}
}
