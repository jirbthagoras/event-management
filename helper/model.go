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

func ToEventResponse(event *model.Event) *web.EventResponse {
	return &web.EventResponse{
		Id:          event.Id,
		Name:        event.Name,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
	}
}

func ToEventResponses(events []*model.Event) []*web.EventResponse {
	var categoryResponses []*web.EventResponse

	for _, event := range events {
		categoryResponses = append(categoryResponses, ToEventResponse(event))
	}

	return categoryResponses
}

func CreateWebResponse(status string, data interface{}) *web.GlobalResponse {
	return &web.GlobalResponse{
		Status: status,
		Data:   data,
	}
}
