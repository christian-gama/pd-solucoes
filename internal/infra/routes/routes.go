package routes

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/internal/infra/router/routing"
	"github.com/christian-gama/pd-solucoes/internal/presentation/middleware"
)

func Global() *routing.Routing {
	return &routing.Routing{
		Middlewares: []http.Middleware{
			middleware.MakeError(),
		},
	}
}

func College() *routing.Routing {
	return &routing.Routing{
		Group:  "/college",
		Routes: []*routing.Route{},
	}
}
