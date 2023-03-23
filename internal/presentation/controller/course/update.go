package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateCourse is a controller to update a course.
type UpdateCourse = http.Controller

// NewUpdateCourse returns a new controller to update a course.
func NewUpdateCourse(s service.UpdateCourse) UpdateCourse {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.UpdateCourseInput) {
			course, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, course)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
