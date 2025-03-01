package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/exception"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/repository"
)

type TicketService interface {
	Create(ctx context.Context, request *web.TicketRequest) *web.TicketResponse
	UpdateStatus(ctx context.Context, id int) *web.TicketResponse
	FindById(ctx context.Context, id int) *web.TicketResponse
	FindAll(ctx context.Context) []*web.TicketResponse
}

type TicketServiceImpl struct {
	TicketRepository   repository.TicketRepository
	AttendeeRepository repository.AttendeeRepository
	EventRepository    repository.EventRepository
	Validator          *validator.Validate
	DB                 *sql.DB
}

func NewTicketServiceImpl(ticketRepository repository.TicketRepository, attendeeRepository repository.AttendeeRepository, eventRepository repository.EventRepository, validator *validator.Validate, DB *sql.DB) *TicketServiceImpl {
	return &TicketServiceImpl{TicketRepository: ticketRepository, AttendeeRepository: attendeeRepository, EventRepository: eventRepository, Validator: validator, DB: DB}
}

func (service TicketServiceImpl) Create(ctx context.Context, request *web.TicketRequest) *web.TicketResponse {

	err := service.Validator.StructCtx(ctx, request)
	exception.PanicIfValidationErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ticket := &model.Ticket{
		Event:    &model.Event{},
		Attendee: &model.Attendee{},
	}

	ticket.Event.Id = request.EventId

	_, err = service.EventRepository.FindById(ctx, tx, ticket.Event.Id)
	helper.PanicIfError(err)

	ticket.Attendee.Id = request.AttendeeId

	_, err = service.AttendeeRepository.FindById(ctx, tx, ticket.Attendee.Id)
	helper.PanicIfError(err)

	ticket, err = service.TicketRepository.Save(ctx, tx, ticket)
	helper.PanicIfError(err)

	ticket, err = service.TicketRepository.FindById(ctx, tx, ticket.Id)
	helper.PanicIfError(err)

	return helper.ToTicketResponse(ticket)
}

func (service TicketServiceImpl) UpdateStatus(ctx context.Context, id int) *web.TicketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ticket, err := service.TicketRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	ticket, err = service.TicketRepository.Update(ctx, tx, &model.Ticket{Id: id})
	helper.PanicIfError(err)

	return helper.ToTicketResponse(ticket)
}

func (service TicketServiceImpl) FindById(ctx context.Context, id int) *web.TicketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ticket, err := service.TicketRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToTicketResponse(ticket)
}

func (service TicketServiceImpl) FindAll(ctx context.Context) []*web.TicketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tickets, err := service.TicketRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToTicketResponses(tickets)
}
