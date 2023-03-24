package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindOneCourseEnrollment is a controller to find one courseEnrollment.
type FindOneCourseEnrollment = http.Controller

// NewFindOneCourseEnrollment returns a new controller to find one courseEnrollment.
func NewFindOneCourseEnrollment(s service.FindOneCourseEnrollment) FindOneCourseEnrollment {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindOneCourseEnrollmentInput) {
			courseEnrollment, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, courseEnrollment)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
			Params: []string{"id"},
		},
	)
}
