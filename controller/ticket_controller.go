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
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Cancel(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}

type TicketControllerImpl struct {
	TicketService service.TicketService
}

func NewTicketControllerImpl(ticketService service.TicketService) *TicketControllerImpl {
	return &TicketControllerImpl{TicketService: ticketService}
}

func (controller TicketControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	var rawRequest map[string]string
	helper.ReadFromRequestBody(request, &rawRequest)

	eventId, err := strconv.Atoi(rawRequest["event_id"])
	if err != nil {
		helper.WriteResponseToBody(writer, http.StatusBadRequest, helper.CreateWebResponse("INVALID EVENT ID", nil))
		return
	}

	attendeeId, err := strconv.Atoi(rawRequest["attendee_id"])
	if err != nil {
		helper.WriteResponseToBody(writer, http.StatusBadRequest, helper.CreateWebResponse("INVALID ATTENDEE ID", nil))
		return
	}

	ticketRequest := &web.TicketRequest{
		EventId:    eventId,
		AttendeeId: attendeeId,
	}

	ticketResponse := controller.TicketService.Create(request.Context(), ticketRequest)

	helper.WriteResponseToBody(writer, http.StatusCreated, helper.CreateWebResponse("CREATED", ticketResponse))

}

func (controller TicketControllerImpl) Cancel(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("ticketId")
	ticketId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	_ = controller.TicketService.UpdateStatus(request.Context(), ticketId)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("CANCELED", nil))
}

func (controller TicketControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("ticketId")
	ticketId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	ticketResponse := controller.TicketService.FindById(request.Context(), ticketId)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", ticketResponse))
}

func (controller TicketControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	ticketResponse := controller.TicketService.FindAll(request.Context())

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", ticketResponse))
}
