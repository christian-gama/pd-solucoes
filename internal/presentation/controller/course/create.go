package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateCourse is a controller to create a course.
type CreateCourse = http.Controller

// NewCreateCourse returns a new controller to create a course.
func NewCreateCourse(s service.CreateCourse) CreateCourse {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.CreateCourseInput) {
			course, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, course)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPost,
		},
	)
}
