//go:build wireinjector
// +build wireinjector

package main

import (
	"github.com/google/wire"
	"jirbthagoras/event-management/app"
	"jirbthagoras/event-management/controller"
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

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		app.NewValidator,
		app.NewServer,
		adminControllerSet,
		app.NewControllers,
		app.NewRouter,
	)

	return nil
}
