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
