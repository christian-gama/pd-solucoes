package service

type DeleteCourseEnrollmentInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
