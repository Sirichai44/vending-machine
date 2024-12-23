package dtos

type ErrorValidation struct {
	Tag     string `json:"tag"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
