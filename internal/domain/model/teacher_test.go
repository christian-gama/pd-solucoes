package model_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type TeacherSuite struct {
	suite.Suite
}

func TestTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(TeacherSuite))
}

func (s *TeacherSuite) TestNewTeacher() {
	type Sut struct {
		Sut  func() (*model.Teacher, error)
		Data *model.Teacher
	}

	makeSut := func() *Sut {
		data := fake.Teacher()

		sut := func() (*model.Teacher, error) {
			return model.NewTeacher(data.ID, data.Name, data.Degree)
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

	s.Run("should return an error when 'degree' is empty", func() {
		sut := makeSut()

		sut.Data.Degree = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
