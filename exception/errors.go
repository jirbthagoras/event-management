package exception

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AppError interface {
	Error() string
	GetStatus() string
	GetCode() int
}

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) GetCode() int {
	return http.StatusNotFound
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

func (e *NotFoundError) GetStatus() string {
	return "NOT FOUND"
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{Msg: msg}
}

type ValidationError struct {
	Errors map[string]interface{}
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		Errors: make(map[string]interface{}),
	}
}

func PanicIfValidationErr(err error) {
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationError := NewValidationError()
			for _, fieldError := range validationErrors {
				validationError.Errors[fieldError.Field()] = fmt.Sprintf(
					"Error on field '%s': '%v' failed rule '%s'",
					fieldError.Field(),
					fieldError.Value(),
					fieldError.Tag(),
				)
			}
			panic(validationError)
		}
		panic(err)
	}
}

func (v ValidationError) Error() string {
	return "failed validation"
}

func (v ValidationError) GetStatus() string {
	return "BAD REQUEST"
}

func (v ValidationError) GetCode() int {
	return http.StatusBadRequest
}
