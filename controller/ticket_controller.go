package controller

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/service"
	"net/http"
	"strconv"
)

type TicketController interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Cancel(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request)
}

type TicketControllerImpl struct {
	TicketService service.TicketService
}

func NewTicketControllerImpl(ticketService service.TicketService) *TicketControllerImpl {
	return &TicketControllerImpl{TicketService: ticketService}
}

func (controller TicketControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	ticketRequest := &web.TicketRequest{}
	helper.ReadFromRequestBody(request, ticketRequest)

	ticketResponse := controller.TicketService.Create(request.Context(), ticketRequest)

	helper.WriteResponseToBody(writer, http.StatusCreated, helper.CreateWebResponse("CREATED", ticketResponse))
}

func (controller TicketControllerImpl) Cancel(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("ticketId")
	ticketId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	ticketResponse := controller.TicketService.UpdateStatus(request.Context(), ticketId)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("CANCELED", ticketResponse))
}

func (controller TicketControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("ticketId")
	ticketId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	ticketResponse := controller.TicketService.FindById(request.Context(), ticketId)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", ticketResponse))
}

func (controller TicketControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	ticketResponse := controller.TicketService.FindAll(request.Context())

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", ticketResponse))
}
