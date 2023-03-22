package service

type UpdateCollegeInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateCollegeInput
}

type UpdateCollegeOutput = FindOneCollegeOutput
