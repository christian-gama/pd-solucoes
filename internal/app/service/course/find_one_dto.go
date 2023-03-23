package service

type FindOneCourseInput struct {
	ID uint `validate:"required" uri:"id"`
}
