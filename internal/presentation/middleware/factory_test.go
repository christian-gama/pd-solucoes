package middleware_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/presentation/middleware"
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
		middleware.MakeError()
	})
}
