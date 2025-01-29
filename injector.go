//go:build wireinjector
// +build wireinjector

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

var adminControllerSet = wire.NewSet(
	repository.NewAdminRepositoryImpl,
	wire.Bind(new(repository.AdminRepository), new(*repository.AdminRepositoryImpl)),
	service.NewAdminServiceImpl,
	wire.Bind(new(service.AdminService), new(*service.AdminServiceImpl)),
	controller.NewAdminControllerImpl,
	wire.Bind(new(controller.AdminController), new(*controller.AdminControllerImpl)),
)

var eventControllerSet = wire.NewSet(
	repository.NewEventRepositoryImpl,
	wire.Bind(new(repository.EventRepository), new(*repository.EventRepositoryImpl)),
	service.NewEventServiceImpl,
	wire.Bind(new(service.EventService), new(*service.EventServiceImpl)),
	controller.NewEventControllerImpl,
	wire.Bind(new(controller.EventController), new(*controller.EventControllerImpl)),
)

var attendeeControllerSet = wire.NewSet(
	repository.NewAttendeeRepositoryImpl,
	wire.Bind(new(repository.AttendeeRepository), new(*repository.AttendeeRepositoryImpl)),
	service.NewAttendeeServiceImpl,
	wire.Bind(new(service.AttendeeService), new(*service.AttendeeServiceImpl)),
	controller.NewAttendeeControllerImpl,
	wire.Bind(new(controller.AttendeeController), new(*controller.AttendeeControllerImpl)),
)

var ticketControllerSet = wire.NewSet(
	repository.NewTicketRepositoryImpl,
	wire.Bind(new(repository.TicketRepository), new(*repository.TicketRepositoryImpl)),
	service.NewTicketServiceImpl,
	wire.Bind(new(service.TicketService), new(*service.TicketServiceImpl)),
	controller.NewTicketControllerImpl,
	wire.Bind(new(controller.TicketController), new(*controller.TicketControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		app.NewValidator,
		app.NewServer,
		adminControllerSet,
		eventControllerSet,
		attendeeControllerSet,
		ticketControllerSet,
		middleware.NewAuthMiddleware,
		app.NewControllers,
		app.NewRouter,
	)

	return nil
}
