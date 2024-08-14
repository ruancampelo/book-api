package models

type ObjectResponse struct {
	IsValid bool        `json:"is_valid"`
	Error   string      `json:"error"`
	Object  interface{} `json:"object"`
}

func CreateResultOk(o interface{}) ObjectResponse {
	return ObjectResponse{
		IsValid: true,
		Object:  o,
	}
}

func CreateResultErro(erro string) ObjectResponse {
	return ObjectResponse{
		IsValid: false,
		Error:   erro,
	}
}
