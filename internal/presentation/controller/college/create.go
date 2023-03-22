package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateCollege is a controller to create a college.
type CreateCollege = http.Controller

// NewCreateCollege returns a new controller to create a college.
func NewCreateCollege(s service.CreateCollege) CreateCollege {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.CreateCollegeInput) {
			college, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, college)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPost,
		},
	)
}
