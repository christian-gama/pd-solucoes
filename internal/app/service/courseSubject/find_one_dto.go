package service

type FindOneCourseSubjectInput struct {
	ID uint `validate:"required" uri:"id" faker:"uint"`
}
