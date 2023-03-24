package service

type UpdateInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateInput
}
