package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateTeacher is a controller to update a teacher.
type UpdateTeacher = http.Controller

// NewUpdateTeacher returns a new controller to update a teacher.
func NewUpdateTeacher(s service.UpdateTeacher) UpdateTeacher {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.UpdateTeacherInput) {
			teacher, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, teacher)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
