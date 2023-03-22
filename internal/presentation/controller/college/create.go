package controller

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// CreateCollege is a controller to create a college.
type CreateCollege = http.Controller

// NewCreateCollege creates a new controller to create a college.
func NewCreateCollege(service service.CreateCollege) CreateCollege {
	if service == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *dto.CreateCollegeInput) {
			college, err := service.Handle(ctx.Request.Context(), input)
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
