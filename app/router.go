package app

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/controller"
	"jirbthagoras/event-management/exception"
)

type Controllers struct {
	controller.AdminController
	controller.EventController
}

func NewControllers(adminController controller.AdminController, eventController controller.EventController) *Controllers {
	return &Controllers{AdminController: adminController, EventController: eventController}
}

func NewRouter(controllers *Controllers) *httprouter.Router {
	router := httprouter.New()
	registerAdminRoute(router, controllers.AdminController)
	registerEventRoute(router, controllers.EventController)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func registerAdminRoute(router *httprouter.Router, adminController controller.AdminController) {
	router.POST("/api/admin/login", adminController.Login)
}

func registerEventRoute(router *httprouter.Router, eventController controller.EventController) {
	router.POST("/api/event", eventController.Create)
	router.GET("/api/event", eventController.FindAll)
	router.PUT("/api/event/:categoryId", eventController.Update)
	router.GET("/api/event/:categoryId", eventController.FindById)
	router.DELETE("/api/event/:categoryId", eventController.Delete)
}
