package service

import service "github.com/christian-gama/pd-solucoes/internal/app/service/college"

type FindOneCourseInput struct {
	ID uint `validate:"required" uri:"id"`
}

type FindOneCourseOutput struct {
	ID      uint                          `json:"id"`
	Name    string                        `json:"name"`
	College *service.FindOneCollegeOutput `json:"college"`
}
