package service

type FindOneTeacherInput struct {
	ID uint `validate:"required" uri:"id"`
}

type FindOneTeacherOutput struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Degree string `json:"degree"`
}
