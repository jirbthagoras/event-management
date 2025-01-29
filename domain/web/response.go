package web

import (
	"jirbthagoras/event-management/domain/model"
	"time"
)

type AdminResponse struct {
	Token string `json:"token"`
}

type GlobalResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

type EventResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

type AttendeeResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TicketResponse struct {
	Id       int             `json:"id"`
	Event    *model.Event    `json:"event"`
	Attendee *model.Attendee `json:"attendee"`
	Status   string          `json:"status"`
}
