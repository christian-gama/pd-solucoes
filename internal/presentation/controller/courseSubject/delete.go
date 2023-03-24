package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteCourseSubject is a controller to find one courseSubject.
type DeleteCourseSubject = http.Controller

// NewDeleteCourseSubject returns a new controller to find one courseSubject.
func NewDeleteCourseSubject(s service.DeleteCourseSubject) DeleteCourseSubject {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.DeleteCourseSubjectInput) {
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
