package log_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/pkg/log"
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
		log.MakeLog()
		log.MakeLogWithCaller(1)
		log.MakeLogWithStack()
	})
}
