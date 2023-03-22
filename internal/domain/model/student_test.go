package model_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type StudentSuite struct {
	suite.Suite
}

func TestStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(StudentSuite))
}

func (s *StudentSuite) TestNewStudent() {
	type Sut struct {
		Sut  func() (*model.Student, error)
		Data *model.Student
	}

	makeSut := func() *Sut {
		data := fake.Student()

		sut := func() (*model.Student, error) {
			return model.NewStudent(data.ID, data.Name, data.Cpf)
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

	s.Run("should return an error when 'cpf' is empty", func() {
		sut := makeSut()

		sut.Data.Cpf = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
