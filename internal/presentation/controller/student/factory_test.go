package controller_test

import (
	"testing"

	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/student"
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
		controller.MakeCreateStudent()
		controller.MakeFindAllStudents()
		controller.MakeFindOneStudent()
		controller.MakeUpdateStudent()
		controller.MakeDeleteStudent()
	})
}
