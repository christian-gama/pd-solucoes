package service

type CreateTeacherInput struct {
	Name   string `json:"name"   validate:"required,max=100,min=2" faker:"len=50"`
	Degree string `json:"degree" validate:"required,min=2,max=100" faker:"len=50"`
}

type CreateTeacherOutput = FindOneTeacherOutput
