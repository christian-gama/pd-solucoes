package service

type DeleteInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
