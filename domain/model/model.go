package model

import "time"

type Admin struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Event struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

type Attendee struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Ticket struct {
	Id      int    `json:"id"`
	EventId int    `json:"event_id"`
	UserId  int    `json:"user_id"`
	Status  string `json:"status"`
}
