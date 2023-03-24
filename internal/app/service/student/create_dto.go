package service

type CreateInput struct {
	Name string `json:"name" validate:"required,max=100,min=2" faker:"len=50"`
	Cpf  string `json:"cpf"  validate:"required,cpf"           faker:"cpf"`
}
