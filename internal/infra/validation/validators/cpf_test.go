package validators_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/validation/validators"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CpfSuite struct {
	suite.Suite
}

func TestCpfSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CpfSuite))
}

func (s *CpfSuite) TestCpf() {
	s.Run("returns true if the cpf is valid", func() {
		s.True(validators.Cpf("43695100079"))
		s.True(validators.Cpf("16931829041"))
		s.True(validators.Cpf("34383737005"))
	})

	s.Run("returns false if the cpf is invalid", func() {
		s.False(validators.Cpf("43695100078"))
		s.False(validators.Cpf("16931829042"))
		s.False(validators.Cpf("34383737006"))
		s.False(validators.Cpf("4369510007"))
		s.False(validators.Cpf("1693182904"))
		s.False(validators.Cpf("3438373700"))
		s.False(validators.Cpf("4369510007a"))
		s.False(validators.Cpf("1693182904A"))
		s.False(validators.Cpf("3438373700_"))
		s.False(validators.Cpf("4369510007."))
		s.False(validators.Cpf("1693182904 "))
		s.False(validators.Cpf("3438373700\t"))
	})

	s.Run("returns false if the cpf is blacklisted", func() {
		s.False(validators.Cpf("00000000000"))
		s.False(validators.Cpf("11111111111"))
		s.False(validators.Cpf("22222222222"))
		s.False(validators.Cpf("33333333333"))
		s.False(validators.Cpf("44444444444"))
		s.False(validators.Cpf("55555555555"))
		s.False(validators.Cpf("66666666666"))
		s.False(validators.Cpf("77777777777"))
		s.False(validators.Cpf("88888888888"))
		s.False(validators.Cpf("99999999999"))
		s.False(validators.Cpf("12345678909"))
	})
}
