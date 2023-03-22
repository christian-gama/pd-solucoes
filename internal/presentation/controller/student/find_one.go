package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindOneStudent is a controller to find one student.
type FindOneStudent = http.Controller

// NewFindOneStudent returns a new controller to find one student.
func NewFindOneStudent(s service.FindOneStudent) FindOneStudent {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindOneStudentInput) {
			student, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, student)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
			Params: []string{"id"},
		},
	)
}
