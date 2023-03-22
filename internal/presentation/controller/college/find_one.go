package controller

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindOneCollege is a controller to findOne college.
type FindOneCollege = http.Controller

// NewFindOneCollege returns a new controller to find one college.
func NewFindOneCollege(service service.FindOneCollege) FindOneCollege {
	if service == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *dto.FindOneCollegeInput) {
			college, err := service.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, college)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
			Params: []string{"id"},
		},
	)
}
