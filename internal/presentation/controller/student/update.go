package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateStudent is a controller to update a student.
type UpdateStudent = http.Controller

// NewUpdateStudent returns a new controller to update a student.
func NewUpdateStudent(s service.UpdateStudent) UpdateStudent {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.UpdateStudentInput) {
			student, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, student)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
