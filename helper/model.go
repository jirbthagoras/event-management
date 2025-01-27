package helper

import (
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/domain/web"
)

func ToAdminResponse(admin model.Admin) *web.AdminResponse {
	return &web.AdminResponse{
		Token: admin.Token,
	}
}

func CreateWebResponse(code int, status string, data interface{}) *web.GlobalResponse {
	return &web.GlobalResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
