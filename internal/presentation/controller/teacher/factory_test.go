package controller_test

import (
	"testing"

	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/teacher"
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
		controller.MakeCreateTeacher()
		controller.MakeFindAllTeachers()
		controller.MakeFindOneTeacher()
		controller.MakeUpdateTeacher()
		controller.MakeDeleteTeacher()
	})
}
