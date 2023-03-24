package service

type CreateInput struct {
	CourseID  uint `json:"courseID"  validate:"required,numeric" faker:"uint"`
	SubjectID uint `json:"subjectID" validate:"required,numeric" faker:"uint"`
}

type CreateOutput struct {
	ID        uint `json:"id"`
	CourseID  uint `json:"courseID"`
	SubjectID uint `json:"subjectID"`
}
