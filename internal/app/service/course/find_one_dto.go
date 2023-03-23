package service

type FindOneCourseInput struct {
	ID uint `validate:"required" uri:"id"`
}

type FindOneCourseOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
