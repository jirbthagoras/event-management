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
}

func NewControllers(adminController controller.AdminController, eventController controller.EventController) *Controllers {
	return &Controllers{AdminController: adminController, EventController: eventController}
}

func NewRouter(controllers *Controllers, middleware *middleware.AuthMiddleware) *httprouter.Router {
	router := httprouter.New()
	registerAdminRoute(router, controllers.AdminController)
	registerEventRoute(router, controllers.EventController, middleware)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func registerAdminRoute(router *httprouter.Router, adminController controller.AdminController) {
	router.POST("/api/admin/login", adminController.Login)
}

func registerEventRoute(router *httprouter.Router, eventController controller.EventController, middleware *middleware.AuthMiddleware) {
	router.POST("/api/event", middleware.Handle(eventController.Create))
	router.GET("/api/event", middleware.Handle(eventController.FindAll))
	router.PUT("/api/event/:categoryId", middleware.Handle(eventController.Update))
	router.GET("/api/event/:categoryId", middleware.Handle(eventController.FindById))
	router.DELETE("/api/event/:categoryId", middleware.Handle(eventController.Delete))
}
