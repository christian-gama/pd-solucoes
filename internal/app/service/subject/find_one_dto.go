package service

type FindOneInput struct {
	ID uint `validate:"required" uri:"id" faker:"uint"`
}
