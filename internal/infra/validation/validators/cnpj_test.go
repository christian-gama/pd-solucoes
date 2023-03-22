package validators_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/validation/validators"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CnpjSuite struct {
	suite.Suite
}

func TestCnpjSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CnpjSuite))
}

func (s *CnpjSuite) TestCnpj() {
	s.Run("returns true if the CNPJ is valid", func() {
		s.True(validators.Cnpj("65306063201790"))
		s.True(validators.Cnpj("02867157457519"))
		s.True(validators.Cnpj("01758561333272"))
	})

	s.Run("returns false if the CNPJ is invalid", func() {
		s.False(validators.Cnpj("43695100078"))
		s.False(validators.Cnpj("16931829042"))
		s.False(validators.Cnpj("34383737006"))
		s.False(validators.Cnpj("4369510007"))
		s.False(validators.Cnpj("1693182904"))
		s.False(validators.Cnpj("3438373700"))
		s.False(validators.Cnpj("4369510007a"))
		s.False(validators.Cnpj("1693182904A"))
		s.False(validators.Cnpj("3438373700_"))
		s.False(validators.Cnpj("4369510007."))
		s.False(validators.Cnpj("1693182904 "))
		s.False(validators.Cnpj("3438373700\t"))
	})
}
