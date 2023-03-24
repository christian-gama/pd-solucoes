package service

type CreateInput struct {
	Name   string `json:"name"   validate:"required,max=100,min=2" faker:"len=50"`
	Degree string `json:"degree" validate:"required,min=2,max=100" faker:"len=50"`
}

type CreateOutput struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Degree string `json:"degree"`
}
