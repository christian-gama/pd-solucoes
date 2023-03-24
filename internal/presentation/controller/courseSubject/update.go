package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateCourseSubject is a controller to update a courseSubject.
type UpdateCourseSubject = http.Controller

// NewUpdateCourseSubject returns a new controller to update a courseSubject.
func NewUpdateCourseSubject(s service.UpdateCourseSubject) UpdateCourseSubject {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.UpdateCourseSubjectInput) {
			courseSubject, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, courseSubject)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
