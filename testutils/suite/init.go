package suite

import (
	"os"

	"github.com/christian-gama/pd-solucoes/internal/infra/env"
	"github.com/christian-gama/pd-solucoes/testutils/faker"
)

func init() {
	faker.InitializeProviders()
	faker.Setup()

	env.Load(".env.test")
	docker := os.Getenv("DOCKER")
	if docker == "true" {
		env.DB.Host = "psql_test"
	}
}
