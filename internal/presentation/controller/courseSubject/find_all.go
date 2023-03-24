package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindAllCourseSubjects is a controller to find all courseSubjects.
type FindAllCourseSubjects = http.Controller

// NewFindAllCourseSubjects returns a new controller to find all courseSubjects.
func NewFindAllCourseSubjects(s service.FindAllCourseSubjects) FindAllCourseSubjects {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindAllInput) {
			courseSubject, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, courseSubject)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
		},
	)
}
