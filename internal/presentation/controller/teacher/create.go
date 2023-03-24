package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateTeacher is a controller to create a teacher.
type CreateTeacher = http.Controller

// NewCreateTeacher returns a new controller to create a teacher.
func NewCreateTeacher(s service.CreateTeacher) CreateTeacher {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.CreateInput) {
			teacher, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, teacher)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPost,
		},
	)
}
