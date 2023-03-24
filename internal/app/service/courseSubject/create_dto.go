package service

type CreateCourseSubjectInput struct {
	CourseID  uint `json:"courseID"  validate:"required,numeric" faker:"uint"`
	SubjectID uint `json:"subjectID" validate:"required,numeric" faker:"uint"`
}
