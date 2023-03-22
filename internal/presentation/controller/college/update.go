package controller

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateCollege is a controller to update a college.
type UpdateCollege = http.Controller

// NewUpdateCollege returns a new controller to update a college.
func NewUpdateCollege(service service.UpdateCollege) UpdateCollege {
	if service == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *dto.UpdateCollegeInput) {
			college, err := service.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, college)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
