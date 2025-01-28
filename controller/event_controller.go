package controller

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/service"
	"net/http"
)

type EventController interface {
	Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}

type EventControllerImpl struct {
	CategoryService service.EventService
}

func (controller EventControllerImpl) Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	var eventRegisterRequest *web.EventRequest
	helper.ReadFromRequestBody(request, eventRegisterRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), eventRegisterRequest)

	helper.WriteResponseToBody(writer, http.StatusCreated, helper.CreateWebResponse("SUCCESSFULLY CREATED", categoryResponse))
}

func (controller EventControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller EventControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller EventControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller EventControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
