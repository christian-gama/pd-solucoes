package service

type UpdateCourseInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateInput
}

type UpdateOutput = CreateOutput
