package service

type DeleteTeacherInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
