package service

type DeleleInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
