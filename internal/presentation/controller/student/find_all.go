package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindAllStudents is a controller to find all students.
type FindAllStudents = http.Controller

// NewFindAllStudents returns a new controller to find all students.
func NewFindAllStudents(s service.FindAllStudents) FindAllStudents {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindAllStudentsInput) {
			student, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, student)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
		},
	)
}
