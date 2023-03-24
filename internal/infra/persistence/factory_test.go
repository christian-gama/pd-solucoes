package persistence_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
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
		persistence.MakeCollege()
		persistence.MakeCourse()
		persistence.MakeCourseEnrollment()
		persistence.MakeCourseSubject()
		persistence.MakeStudent()
		persistence.MakeSubject()
		persistence.MakeTeacher()
	})
}
