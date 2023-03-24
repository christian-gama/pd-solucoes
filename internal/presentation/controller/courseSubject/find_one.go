package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindOneCourseSubject is a controller to find one courseSubject.
type FindOneCourseSubject = http.Controller

// NewFindOneCourseSubject returns a new controller to find one courseSubject.
func NewFindOneCourseSubject(s service.FindOneCourseSubject) FindOneCourseSubject {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindOneCourseSubjectInput) {
			courseSubject, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, courseSubject)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
			Params: []string{"id"},
		},
	)
}
