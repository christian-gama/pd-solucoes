package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateCourseEnrollment is a controller to update a courseEnrollment.
type UpdateCourseEnrollment = http.Controller

// NewUpdateCourseEnrollment returns a new controller to update a courseEnrollment.
func NewUpdateCourseEnrollment(s service.UpdateCourseEnrollment) UpdateCourseEnrollment {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.UpdateInput) {
			courseEnrollment, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, courseEnrollment)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
