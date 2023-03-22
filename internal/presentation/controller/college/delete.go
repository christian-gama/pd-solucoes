package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteCollege is a controller to find one college.
type DeleteCollege = http.Controller

// NewDeleteCollege returns a new controller to find one college.
func NewDeleteCollege(s service.DeleteCollege) DeleteCollege {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.DeleteCollegeInput) {
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
