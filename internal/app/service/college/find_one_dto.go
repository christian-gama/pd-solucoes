package service

type FindOneCollegeInput struct {
	ID uint `validate:"required" uri:"id"`
}

type FindOneCollegeOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}
