package app

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/controller"
)

type Controllers struct {
	controller.AdminController
}

func NewControllers(adminController controller.AdminController) *Controllers {
	return &Controllers{AdminController: adminController}
}

func NewRouter(controllers *Controllers) *httprouter.Router {
	router := httprouter.New()
	registerAdminRoute(router, controllers.AdminController)

	return router
}

func registerAdminRoute(router *httprouter.Router, adminController controller.AdminController) {
	router.POST("/api/admin/login", adminController.Login)
}
