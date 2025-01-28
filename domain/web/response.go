package web

type AdminResponse struct {
	Token string `json:"token"`
}

type GlobalResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}
