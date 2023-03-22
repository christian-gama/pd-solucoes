package suite

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/testutils"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) Skip(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skip()
	})
}

func (s *Suite) Todo(name string, f func()) bool {
	return s.Run(name, func() {
		s.T().Skipf("TODO: %s", name)
	})
}

type SuiteWithConn struct {
	Suite
}

func TestSetupTestsSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(SuiteWithConn))
}

func (s *SuiteWithConn) Run(name string, f func(tx *gorm.DB)) bool {
	return s.Suite.Run(name, func() {
		testutils.Transaction(func(tx *gorm.DB) {
			f(tx)
		})
	})
}
