package service

type FindOneStudentInput struct {
	ID uint `validate:"required" uri:"id" faker:"uint"`
}
