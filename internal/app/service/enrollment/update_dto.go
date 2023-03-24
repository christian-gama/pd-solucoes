package service

type UpdateCourseEnrollmentInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateCourseEnrollmentInput
}
