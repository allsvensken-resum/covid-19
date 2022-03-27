package errs

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError() error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: "Data Not Found",
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}
