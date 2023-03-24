package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteCourseEnrollment is a controller to find one courseEnrollment.
type DeleteCourseEnrollment = http.Controller

// NewDeleteCourseEnrollment returns a new controller to find one courseEnrollment.
func NewDeleteCourseEnrollment(s service.DeleteCourseEnrollment) DeleteCourseEnrollment {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.DeleteInput) {
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
