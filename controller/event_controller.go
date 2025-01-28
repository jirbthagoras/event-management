package controller

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/service"
	"net/http"
	"strconv"
)

type EventController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}

type EventControllerImpl struct {
	EventService service.EventService
}

func NewEventControllerImpl(eventService service.EventService) *EventControllerImpl {
	return &EventControllerImpl{EventService: eventService}
}

func (controller EventControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	eventCreateRequest := &web.EventRequest{}
	helper.ReadFromRequestBody(request, eventCreateRequest)

	eventResponse := controller.EventService.Create(request.Context(), eventCreateRequest)

	helper.WriteResponseToBody(writer, http.StatusCreated, helper.CreateWebResponse("SUCCESSFULLY CREATED", eventResponse))
}

func (controller EventControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	eventUpdateRequest := &web.EventRequest{}
	helper.ReadFromRequestBody(request, eventUpdateRequest)
	eventUpdateRequest.Id = id

	eventResponse := controller.EventService.Update(request.Context(), eventUpdateRequest)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESSFULLY UPDATED", eventResponse))
}

func (controller EventControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	eventResponse := controller.EventService.FindById(request.Context(), id)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", eventResponse))
}

func (controller EventControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	eventResponse := controller.EventService.FindAll(request.Context())

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", eventResponse))
}

func (controller EventControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	categoryId := param.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.EventService.DeleteById(request.Context(), id)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("DELETION SUCCESS", nil))
}
