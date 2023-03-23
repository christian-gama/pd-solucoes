package service

type UpdateTeacherInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateTeacherInput
}
