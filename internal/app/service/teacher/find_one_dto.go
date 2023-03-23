package service

type FindOneTeacherInput struct {
	ID uint `validate:"required" uri:"id"`
}
