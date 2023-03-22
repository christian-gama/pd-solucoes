package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteTeacher is a controller to find one teacher.
type DeleteTeacher = http.Controller

// NewDeleteTeacher returns a new controller to find one teacher.
func NewDeleteTeacher(s service.DeleteTeacher) DeleteTeacher {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.DeleteTeacherInput) {
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
