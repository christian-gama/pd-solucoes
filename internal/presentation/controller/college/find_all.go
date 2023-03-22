package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindAllColleges is a controller to find all colleges.
type FindAllColleges = http.Controller

// NewFindAllColleges returns a new controller to find all colleges.
func NewFindAllColleges(s service.FindAllColleges) FindAllColleges {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindAllCollegesInput) {
			college, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, college)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
		},
	)
}
