package service

type DeleteCourseInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
