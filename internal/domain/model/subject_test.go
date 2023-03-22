package model_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type SubjectSuite struct {
	suite.Suite
}

func TestSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SubjectSuite))
}

func (s *SubjectSuite) TestNewSubject() {
	type Sut struct {
		Sut  func() (*model.Subject, error)
		Data *model.Subject
	}

	makeSut := func() *Sut {
		data := fake.Subject()

		sut := func() (*model.Subject, error) {
			return model.NewSubject(data.ID, data.Name, data.TeacherID, data.Teacher)
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

	s.Run("should return an error when 'teacherID' is zero", func() {
		sut := makeSut()

		sut.Data.TeacherID = 0

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'teacher' is nil", func() {
		sut := makeSut()

		sut.Data.Teacher = nil

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'teacher' is invalid", func() {
		sut := makeSut()

		sut.Data.Teacher.Name = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
