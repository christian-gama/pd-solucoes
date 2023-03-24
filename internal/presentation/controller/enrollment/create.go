package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateCourseEnrollment is a controller to create a courseEnrollment.
type CreateCourseEnrollment = http.Controller

// NewCreateCourseEnrollment returns a new controller to create a courseEnrollment.
func NewCreateCourseEnrollment(s service.CreateCourseEnrollment) CreateCourseEnrollment {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.CreateInput) {
			courseEnrollment, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, courseEnrollment)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPost,
		},
	)
}
