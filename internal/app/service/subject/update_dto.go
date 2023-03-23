package service

type UpdateSubjectInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateSubjectInput
}
