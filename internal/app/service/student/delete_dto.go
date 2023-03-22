package service

type DeleteStudentInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
