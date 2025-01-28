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

func CreateWebResponse(status string, data interface{}) *web.GlobalResponse {
	return &web.GlobalResponse{
		Status: status,
		Data:   data,
	}
}
