package service

type CreateCollegeInput struct {
	Name string `json:"name" validate:"required,max=100,min=2" faker:"len=50"`
	Cnpj string `json:"cnpj" validate:"required,cnpj"          faker:"cnpj"`
}

type CreateOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}
