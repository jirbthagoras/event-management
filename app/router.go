package app

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/controller"
	"jirbthagoras/event-management/exception"
	"jirbthagoras/event-management/middleware"
)

type Controllers struct {
	controller.AdminController
	controller.EventController
	controller.AttendeeController
}

func NewControllers(adminController controller.AdminController, eventController controller.EventController, attendeeController controller.AttendeeController) *Controllers {
	return &Controllers{AdminController: adminController, EventController: eventController, AttendeeController: attendeeController}
}

func NewRouter(controllers *Controllers, middleware *middleware.AuthMiddleware) *httprouter.Router {
	router := httprouter.New()
	registerAdminRoute(router, controllers.AdminController)
	registerEventRoute(router, controllers.EventController, middleware)
	registerAttendeeRoute(router, controllers.AttendeeController, middleware)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func registerAdminRoute(router *httprouter.Router, adminController controller.AdminController) {
	router.POST("/api/admin/login", adminController.Login)
}

func registerEventRoute(router *httprouter.Router, eventController controller.EventController, middleware *middleware.AuthMiddleware) {
	router.POST("/api/event", middleware.Handle(eventController.Create))
	router.GET("/api/event", middleware.Handle(eventController.FindAll))
	router.PUT("/api/event/:eventId", middleware.Handle(eventController.Update))
	router.GET("/api/event/:eventId", middleware.Handle(eventController.FindById))
	router.DELETE("/api/event/:eventId", middleware.Handle(eventController.Delete))
}

func registerAttendeeRoute(router *httprouter.Router, attendeeController controller.AttendeeController, middleware *middleware.AuthMiddleware) {
	router.POST("/api/attendee", middleware.Handle(attendeeController.Create))
	router.GET("/api/attendee", middleware.Handle(attendeeController.FindAll))
	router.PUT("/api/attendee/:attendeeId", middleware.Handle(attendeeController.Update))
	router.GET("/api/attendee/:attendeeId", middleware.Handle(attendeeController.FindById))
	router.DELETE("/api/attendee/:attendeeId", middleware.Handle(attendeeController.Delete))
}
