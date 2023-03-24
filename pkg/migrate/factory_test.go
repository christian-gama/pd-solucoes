package migrate_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/pkg/migrate"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type FactorySuite struct {
	suite.Suite
}

func TestFactorySuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(FactorySuite))
}

func (s *FactorySuite) TestFactory() {
	s.NotPanics(func() {
		migrate.MakeMigrate(false)
		migrate.MakeMigrate(true)
	})
}
