package sql_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
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
		sql.MakePostgres()
	})
}
