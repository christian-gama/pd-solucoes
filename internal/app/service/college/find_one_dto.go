package service

type FindOneCollegeInput struct {
	ID uint `validate:"required" uri:"id"`
}
