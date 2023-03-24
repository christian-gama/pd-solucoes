package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateCourseSubject is a controller to create a courseSubject.
type CreateCourseSubject = http.Controller

// NewCreateCourseSubject returns a new controller to create a courseSubject.
func NewCreateCourseSubject(s service.CreateCourseSubject) CreateCourseSubject {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.CreateInput) {
			courseSubject, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, courseSubject)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPost,
		},
	)
}
