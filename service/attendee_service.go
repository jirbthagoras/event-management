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

type AttendeeService interface {
	Create(ctx context.Context, request *web.AttendeeRequest) *web.AttendeeResponse
	Update(ctx context.Context, request *web.AttendeeRequest) *web.AttendeeResponse
	FindById(ctx context.Context, id int) *web.AttendeeResponse
	FindAll(ctx context.Context) []*web.AttendeeResponse
	DeleteById(ctx context.Context, id int)
}

type AttendeeServiceImpl struct {
	DB                 *sql.DB
	Validator          *validator.Validate
	AttendeeRepository repository.AttendeeRepository
}

func (service AttendeeServiceImpl) Create(ctx context.Context, request *web.AttendeeRequest) *web.AttendeeResponse {
	err := service.Validator.Struct(request)
	exception.PanicIfValidationErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attendee := model.Attendee{
		Name:  request.Name,
		Email: request.Email,
	}

	result, err := service.AttendeeRepository.Save(ctx, tx, &attendee)

	return helper.ToAttendeeResponse(result)
}

func (service AttendeeServiceImpl) Update(ctx context.Context, request *web.AttendeeRequest) *web.AttendeeResponse {
	err := service.Validator.StructCtx(ctx, request)
	exception.PanicIfValidationErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.AttendeeRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	attendee := &model.Attendee{
		Id:    request.Id,
		Name:  request.Name,
		Email: request.Email,
	}

	attendee, err = service.AttendeeRepository.Update(ctx, tx, attendee)
	helper.PanicIfError(err)

	return helper.ToAttendeeResponse(attendee)
}

func (service AttendeeServiceImpl) FindById(ctx context.Context, id int) *web.AttendeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attendee, err := service.AttendeeRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToAttendeeResponse(attendee)
}

func (service AttendeeServiceImpl) FindAll(ctx context.Context) []*web.AttendeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attendee, err := service.AttendeeRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToAttendeeResponses(attendee)
}

func (service AttendeeServiceImpl) DeleteById(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.AttendeeRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	err = service.AttendeeRepository.Delete(ctx, tx, id)
	helper.PanicIfError(err)
}

func NewAttendeeServiceImpl(DB *sql.DB, validator *validator.Validate, attendeeRepository repository.AttendeeRepository) *AttendeeServiceImpl {
	return &AttendeeServiceImpl{DB: DB, Validator: validator, AttendeeRepository: attendeeRepository}
}
