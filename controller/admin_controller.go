package controller

import (
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/service"
	"net/http"
)

type AdminController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func (controller AdminControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminLoginRequest := web.AdminLoginRequest{}
	helper.ReadFromRequestBody(request, &adminLoginRequest)

	var adminResponse *web.AdminResponse = controller.AdminService.Login(request.Context(), &adminLoginRequest)

	helper.WriteResponseToBody(writer, http.StatusOK, helper.CreateWebResponse("SUCCESS", adminResponse))
}

func NewAdminControllerImpl(adminService service.AdminService) *AdminControllerImpl {
	return &AdminControllerImpl{AdminService: adminService}
}
