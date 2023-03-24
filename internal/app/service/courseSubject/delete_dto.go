package service

type DeleteCourseSubjectInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
