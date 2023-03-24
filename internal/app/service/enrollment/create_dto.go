package service

type CreateInput struct {
	StudentID       uint `json:"studentID"       validate:"required,numeric" faker:"uint"`
	CourseSubjectID uint `json:"courseSubjectID" validate:"required,numeric" faker:"uint"`
}

type CreateOutput struct {
	ID              uint `json:"id"`
	StudentID       uint `json:"studentID"`
	CourseSubjectID uint `json:"courseSubjectID"`
}
