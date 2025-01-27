package web

type AdminResponse struct {
	Token string `json:"token"`
}

type GlobalResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
