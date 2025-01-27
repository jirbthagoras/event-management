package web

type AdminResponse struct {
	Token string `json:"token"`
}

type GlobalResponse struct {
	Code   string      `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
