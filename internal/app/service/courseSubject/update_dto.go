package service

type UpdateCourseSubjectInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateCourseSubjectInput
}
