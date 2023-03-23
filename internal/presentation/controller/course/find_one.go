package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindOneCourse is a controller to find one course.
type FindOneCourse = http.Controller

// NewFindOneCourse returns a new controller to find one course.
func NewFindOneCourse(s service.FindOneCourse) FindOneCourse {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindOneCourseInput) {
			course, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, course)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
			Params: []string{"id"},
		},
	)
}
