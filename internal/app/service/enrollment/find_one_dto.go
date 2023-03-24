package service

type FindOneCourseEnrollmentInput struct {
	ID uint `validate:"required" uri:"id"`
}
