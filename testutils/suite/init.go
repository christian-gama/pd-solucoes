package suite

import (
	"github.com/christian-gama/pd-solucoes/testutils/faker"
)

func init() {
	faker.InitializeProviders()
	faker.Setup()
}
