package dtos

type Context struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}