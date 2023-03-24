package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteStudent is a controller to find one student.
type DeleteStudent = http.Controller

// NewDeleteStudent returns a new controller to find one student.
func NewDeleteStudent(s service.DeleteStudent) DeleteStudent {
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
