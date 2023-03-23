package controller

import (
	"errors"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/gin-gonic/gin"
)

// FindAllSubjects is a controller to find all subjects.
type FindAllSubjects = http.Controller

// NewFindAllSubjects returns a new controller to find all subjects.
func NewFindAllSubjects(s service.FindAllSubjects) FindAllSubjects {
	if s == nil {
		panic(errors.New("service cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.FindAllSubjectsInput) {
			subject, err := s.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, subject)
		},

		http.ControllerOptions{
			Path:   "/",
			Method: http.MethodGet,
		},
	)
}
