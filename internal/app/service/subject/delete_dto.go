package service

type DeleteSubjectInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
