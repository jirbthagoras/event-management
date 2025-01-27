package model

import "time"

type Admin struct {
	Id       int
	Email    string
	Password string
	Token    string
}

type Event struct {
	Id          int
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type Attendee struct {
	Id    int
	Name  string
	Email string
}

type Ticket struct {
	Id      int    `json:"id"`
	EventId int    `json:"event_id"`
	UserId  int    `json:"user_id"`
	Status  string `json:"status"`
}
