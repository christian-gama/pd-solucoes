package internal

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/infra/env"
	"github.com/christian-gama/pd-solucoes/internal/infra/router"
	"github.com/christian-gama/pd-solucoes/internal/infra/router/routing"
	"github.com/christian-gama/pd-solucoes/internal/infra/routes"
	"github.com/christian-gama/pd-solucoes/internal/infra/server"
	"github.com/christian-gama/pd-solucoes/pkg/log"

	// Initialize custom validation aliases.
	_ "github.com/christian-gama/pd-solucoes/internal/infra/validation"
)

// Bootstrap is the main function that starts the application.
func Bootstrap(ctx context.Context, log log.Logger, envFile string) {
	env.Load(envFile)
	log.Infof("Booting the application")

	r := router.New()
	routing.Registerer(r.Group("/api/v1"),
		routes.Global(),
		routes.Colleges(),
		routes.Teachers(),
		routes.Students(),
		routes.Courses(),
		routes.Subjects(),
	)

	server.Start(ctx, r, log)
}
