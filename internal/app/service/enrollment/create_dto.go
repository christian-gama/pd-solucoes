package service

type CreateCourseEnrollmentInput struct {
	StudentID       uint `json:"studentID"       validate:"required,numeric" faker:"uint"`
	CourseSubjectID uint `json:"courseSubjectID" validate:"required,numeric" faker:"uint"`
}
