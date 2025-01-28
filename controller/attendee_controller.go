package controller

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/service"
	"net/http"
	"strconv"
)

type AttendeeController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}

type AttendeeControllerImpl struct {
	AttendeeService service.AttendeeService
}

func NewAttendeeControllerImpl(attendeeService service.AttendeeService) *AttendeeControllerImpl {
	return &AttendeeControllerImpl{AttendeeService: attendeeService}
}

func (controller AttendeeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	attendeeCreateRequest := &web.AttendeeRequest{}
	helper.ReadFromRequestBody(request, attendeeCreateRequest)

	attendeeResponse := controller.AttendeeService.Create(request.Context(), attendeeCreateRequest)

	helper.WriteResponseToBody(writer, http.StatusCreated, helper.CreateWebResponse("SUCCESSFULLY CREATED", attendeeResponse))
}

func (controller AttendeeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("attendeeId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	attendeeUpdateRequest := &web.AttendeeRequest{}
	helper.ReadFromRequestBody(request, attendeeUpdateRequest)
	attendeeUpdateRequest.Id = id

	attendeeResponse := controller.AttendeeService.Update(request.Context(), attendeeUpdateRequest)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESSFULLY UPDATED", attendeeResponse))
}

func (controller AttendeeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("attendeeId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	attendeeResponse := controller.AttendeeService.FindById(request.Context(), id)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", attendeeResponse))
}

func (controller AttendeeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	attendeeResponse := controller.AttendeeService.FindAll(request.Context())

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", attendeeResponse))
}

func (controller AttendeeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("attendeeId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.AttendeeService.DeleteById(request.Context(), id)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("DELETION SUCCESS", nil))
}
