package errors

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	Errors struct{ Field map[string][]string } `json:"Errors"`
	Code   int                                 `json:"-"`
}

func NewError(code int, field string, detail string) *HTTPError {
	return &HTTPError{
		Code: code,
		Errors: struct{ Field map[string][]string }{
			Field: map[string][]string{field: {detail}},
		},
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError:%d", e.Code)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return &HTTPError{}
}
