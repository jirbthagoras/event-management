package exception

import "net/http"

type AppError interface {
	Error() string
	GetStatus() string
	GetCode() int
}

type NotFoundError struct {
	Msg    string
	Status string
}

func (e *NotFoundError) GetCode() int {
	return http.StatusNotFound
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

func (e *NotFoundError) GetStatus() string {
	return e.Status
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{Msg: msg, Status: "NOT FOUND"}
}
