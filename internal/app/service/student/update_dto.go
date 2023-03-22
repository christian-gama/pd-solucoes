package service

type UpdateStudentInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateStudentInput
}

type UpdateStudentOutput = FindOneStudentOutput
