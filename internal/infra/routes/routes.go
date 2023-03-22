package routes

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/internal/infra/router/routing"
	collegeController "github.com/christian-gama/pd-solucoes/internal/presentation/controller/college"
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
		Group: "/college",
		Routes: []*routing.Route{
			{Controller: collegeController.MakeCreateCollege()},
			{Controller: collegeController.MakeUpdateCollege()},
			{Controller: collegeController.MakeFindOneCollege()},
		},
	}
}
