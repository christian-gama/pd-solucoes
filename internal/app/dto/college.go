package dto

type CreateCollegeInput struct {
	Name string `json:"name" validate:"required,max=100,min=2"`
	Cnpj string `json:"cnpj" validate:"required"`
}

type CreateCollegeOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}
