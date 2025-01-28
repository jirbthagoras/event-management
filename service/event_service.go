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
	"time"
)

type EventService interface {
	Create(ctx context.Context, request *web.EventRequest) *web.EventResponse
	Update(ctx context.Context, request *web.EventRequest) *web.EventResponse
	FindById(ctx context.Context, request *web.EventRequest) *web.EventResponse
	FindAll(ctx context.Context, request *web.EventRequest) []*web.EventResponse
	DeleteById(ctx context.Context, request *web.EventRequest)
}

type EventServiceImpl struct {
	EventRepository repository.EventRepository
	Validator       *validator.Validate
	DB              *sql.DB
}

func (service EventServiceImpl) Create(ctx context.Context, request *web.EventRequest) *web.EventResponse {
	// Validate the request
	err := service.Validator.StructCtx(ctx, request)
	exception.PanicIfValidationErr(err)

	startTime, err := time.Parse("2006-01-02", request.StartTime)
	helper.PanicIfError(err)

	endTime, err := time.Parse("2006-01-02", request.EndTime)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event := model.Event{
		Name:        request.Name,
		Description: request.Description,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	result, err := service.EventRepository.Save(ctx, tx, &event)

	return helper.ToEventResponse(result)
}

func (service EventServiceImpl) Update(ctx context.Context, request *web.EventRequest) *web.EventResponse {
	//TODO implement me
	panic("implement me")
}

func (service EventServiceImpl) FindById(ctx context.Context, request *web.EventRequest) *web.EventResponse {
	//TODO implement me
	panic("implement me")
}

func (service EventServiceImpl) FindAll(ctx context.Context, request *web.EventRequest) []*web.EventResponse {
	//TODO implement me
	panic("implement me")
}

func (service EventServiceImpl) DeleteById(ctx context.Context, request *web.EventRequest) {
	//TODO implement me
	panic("implement me")
}

func NewEventServiceImpl(eventRepository *repository.EventRepository, validator *validator.Validate, DB *sql.DB) *EventServiceImpl {
	return &EventServiceImpl{EventRepository: eventRepository, Validator: validator, DB: DB}
}
