package service

type FindOneSubjectInput struct {
	ID uint `validate:"required" uri:"id" faker:"uint"`
}
