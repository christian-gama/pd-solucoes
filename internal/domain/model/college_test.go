package model_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CollegeSuite struct {
	suite.Suite
}

func TestCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CollegeSuite))
}

func (s *CollegeSuite) TestNewCollege() {
	type Sut struct {
		Sut  func() (*model.College, error)
		Data *model.College
	}

	makeSut := func() *Sut {
		data := fake.College()

		sut := func() (*model.College, error) {
			return model.NewCollege(data.Name, data.Cnpj)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("should create a new model", func() {
		sut := makeSut()

		model, err := sut.Sut()

		s.NoError(err)
		s.NotNil(model, "model should not be nil")
	})

	s.Run("should return an error when 'name' is empty", func() {
		sut := makeSut()

		sut.Data.Name = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'cnpj' is empty", func() {
		sut := makeSut()

		sut.Data.Cnpj = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
