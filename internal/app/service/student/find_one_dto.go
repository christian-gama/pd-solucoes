package service

type FindOneStudentInput struct {
	ID uint `validate:"required" uri:"id"`
}

type FindOneStudentOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}
