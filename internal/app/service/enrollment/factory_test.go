package service_test

import (
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
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
		service.MakeCreateCourseEnrollment()
		service.MakeFindAllCourseEnrollments()
		service.MakeFindOneCourseEnrollment()
		service.MakeUpdateCourseEnrollment()
		service.MakeDeleteCourseEnrollment()
	})
}
