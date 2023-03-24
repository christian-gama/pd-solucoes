package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// UpdateSubject is a controller to update a subject.
type UpdateSubject = http.Controller

// NewUpdateSubject returns a new controller to update a subject.
func NewUpdateSubject(s service.UpdateSubject) UpdateSubject {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.UpdateInput) {
			subject, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, subject)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodPut,
			Params: []string{"id"},
		},
	)
}
