package service

type CreateSubjectInput struct {
	Name      string `json:"name"      validate:"required,max=100,min=2" faker:"len=50"`
	TeacherID uint   `json:"teacherID" validate:"required,numeric"       faker:"uint"`
}
