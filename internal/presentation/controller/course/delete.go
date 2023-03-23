package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteCourse is a controller to find one course.
type DeleteCourse = http.Controller

// NewDeleteCourse returns a new controller to find one course.
func NewDeleteCourse(s service.DeleteCourse) DeleteCourse {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.DeleteCourseInput) {
			err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.NoContent(ctx)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodDelete,
			Params: []string{"id"},
		},
	)
}
