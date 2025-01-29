package web

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type EventRequest struct {
	Id          int
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	StartTime   string `json:"start_time" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
}

type AttendeeRequest struct {
	Id    int
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type TicketRequest struct {
	EventId    int `json:"event_id" validate:"required"`
	AttendeeId int `json:"attendee_id" validate:"required"`
}
