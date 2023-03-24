package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateStudent is a controller to create a student.
type CreateStudent = http.Controller

// NewCreateStudent returns a new controller to create a student.
func NewCreateStudent(s service.CreateStudent) CreateStudent {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.CreateInput) {
			student, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, student)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPost,
		},
	)
}
