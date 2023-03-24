package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindAllCourseEnrollments is a controller to find all courseEnrollments.
type FindAllCourseEnrollments = http.Controller

// NewFindAllCourseEnrollments returns a new controller to find all courseEnrollments.
func NewFindAllCourseEnrollments(s service.FindAllCourseEnrollments) FindAllCourseEnrollments {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindAllCourseEnrollmentsInput) {
			courseEnrollment, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, courseEnrollment)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
		},
	)
}
