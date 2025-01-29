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
	var eventResponses []*web.EventResponse

	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}

	return eventResponses
}

func ToAttendeeResponse(attendee *model.Attendee) *web.AttendeeResponse {
	return &web.AttendeeResponse{
		Id:    attendee.Id,
		Name:  attendee.Name,
		Email: attendee.Email,
	}
}

func ToAttendeeResponses(attendees []*model.Attendee) []*web.AttendeeResponse {
	var attendeeResponses []*web.AttendeeResponse

	for _, attendee := range attendees {
		attendeeResponses = append(attendeeResponses, ToAttendeeResponse(attendee))
	}

	return attendeeResponses
}

func ToTicketResponse(ticket *model.Ticket) *web.TicketResponse {
	return &web.TicketResponse{
		Id:       ticket.Id,
		Event:    *ticket.Event,
		Attendee: *ticket.Attendee,
		Status:   ticket.Status,
	}
}

func ToTicketResponses(tickets []*model.Ticket) []*web.TicketResponse {
	var ticketResponses []*web.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, ToTicketResponse(ticket))
	}
}

func CreateWebResponse(status string, data interface{}) *web.GlobalResponse {
	return &web.GlobalResponse{
		Status: status,
		Data:   data,
	}
}
