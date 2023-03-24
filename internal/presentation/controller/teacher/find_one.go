package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindOneTeacher is a controller to find one teacher.
type FindOneTeacher = http.Controller

// NewFindOneTeacher returns a new controller to find one teacher.
func NewFindOneTeacher(s service.FindOneTeacher) FindOneTeacher {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindOneInput) {
			teacher, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, teacher)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
			Params: []string{"id"},
		},
	)
}
