package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindAllTeachers is a controller to find all teachers.
type FindAllTeachers = http.Controller

// NewFindAllTeachers returns a new controller to find all teachers.
func NewFindAllTeachers(s service.FindAllTeachers) FindAllTeachers {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindAllInput) {
			teacher, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, teacher)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
		},
	)
}
