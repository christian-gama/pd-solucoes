package service

type DeleteCollegeInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
