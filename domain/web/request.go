package web

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type EventRequest struct {
	Id          int    `json:"id" validate:"required,gte=1"`
	Name        string `json:"category" validate:"required"`
	Description string `json:"description" validate:"required"`
	StartTime   string `json:"start_time" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
}
