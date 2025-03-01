// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"jirbthagoras/event-management/app"
	"jirbthagoras/event-management/controller"
	"jirbthagoras/event-management/middleware"
	"jirbthagoras/event-management/repository"
	"jirbthagoras/event-management/service"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	adminRepositoryImpl := repository.NewAdminRepositoryImpl()
	db := app.NewDB()
	validate := app.NewValidator()
	adminServiceImpl := service.NewAdminServiceImpl(adminRepositoryImpl, db, validate)
	adminControllerImpl := controller.NewAdminControllerImpl(adminServiceImpl)
	eventRepositoryImpl := repository.NewEventRepositoryImpl()
	eventServiceImpl := service.NewEventServiceImpl(eventRepositoryImpl, validate, db)
	eventControllerImpl := controller.NewEventControllerImpl(eventServiceImpl)
	attendeeRepositoryImpl := repository.NewAttendeeRepositoryImpl()
	attendeeServiceImpl := service.NewAttendeeServiceImpl(db, validate, attendeeRepositoryImpl)
	attendeeControllerImpl := controller.NewAttendeeControllerImpl(attendeeServiceImpl)
	ticketRepositoryImpl := repository.NewTicketRepositoryImpl()
	ticketServiceImpl := service.NewTicketServiceImpl(ticketRepositoryImpl, attendeeRepositoryImpl, eventRepositoryImpl, validate, db)
	ticketControllerImpl := controller.NewTicketControllerImpl(ticketServiceImpl)
	controllers := app.NewControllers(adminControllerImpl, eventControllerImpl, attendeeControllerImpl, ticketControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(db, adminRepositoryImpl)
	router := app.NewRouter(controllers, authMiddleware)
	server := app.NewServer(router)
	return server
}

// injector.go:

var adminControllerSet = wire.NewSet(repository.NewAdminRepositoryImpl, wire.Bind(new(repository.AdminRepository), new(*repository.AdminRepositoryImpl)), service.NewAdminServiceImpl, wire.Bind(new(service.AdminService), new(*service.AdminServiceImpl)), controller.NewAdminControllerImpl, wire.Bind(new(controller.AdminController), new(*controller.AdminControllerImpl)))

var eventControllerSet = wire.NewSet(repository.NewEventRepositoryImpl, wire.Bind(new(repository.EventRepository), new(*repository.EventRepositoryImpl)), service.NewEventServiceImpl, wire.Bind(new(service.EventService), new(*service.EventServiceImpl)), controller.NewEventControllerImpl, wire.Bind(new(controller.EventController), new(*controller.EventControllerImpl)))

var attendeeControllerSet = wire.NewSet(repository.NewAttendeeRepositoryImpl, wire.Bind(new(repository.AttendeeRepository), new(*repository.AttendeeRepositoryImpl)), service.NewAttendeeServiceImpl, wire.Bind(new(service.AttendeeService), new(*service.AttendeeServiceImpl)), controller.NewAttendeeControllerImpl, wire.Bind(new(controller.AttendeeController), new(*controller.AttendeeControllerImpl)))

var ticketControllerSet = wire.NewSet(repository.NewTicketRepositoryImpl, wire.Bind(new(repository.TicketRepository), new(*repository.TicketRepositoryImpl)), service.NewTicketServiceImpl, wire.Bind(new(service.TicketService), new(*service.TicketServiceImpl)), controller.NewTicketControllerImpl, wire.Bind(new(controller.TicketController), new(*controller.TicketControllerImpl)))
