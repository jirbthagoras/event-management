package exception

import (
	"errors"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	var validationError *ValidationError
	if errors.As(err.(error), &validationError) {
		launchValidationError(writer, *validationError)
		return
	}

	var appError AppError
	if errors.As(err.(error), &appError) {
		launchCustomError(writer, appError)
		return
	}

	launchInternalError(writer, err.(error))
}

func launchCustomError(writer http.ResponseWriter, err AppError) {
	globalResponse := &web.GlobalResponse{
		Status: err.GetStatus(),
		Errors: err.Error(),
	}
	helper.WriteResponseToBody(writer, err.GetCode(), globalResponse)
}

func launchValidationError(writer http.ResponseWriter, err ValidationError) {
	globalResponse := &web.GlobalResponse{
		Status: err.GetStatus(),
		Errors: err.Errors,
	}
	helper.WriteResponseToBody(writer, err.GetCode(), globalResponse)
}

func launchInternalError(writer http.ResponseWriter, err error) {
	globalResponse := &web.GlobalResponse{
		Status: "INTERNAL SERVER ERROR",
		Errors: err.Error(),
	}
	helper.WriteResponseToBody(writer, http.StatusInternalServerError, globalResponse)
}
