package model_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CourseSuite struct {
	suite.Suite
}

func TestCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CourseSuite))
}

func (s *CourseSuite) TestNewCourse() {
	type Sut struct {
		Sut  func() (*model.Course, error)
		Data *model.Course
	}

	makeSut := func() *Sut {
		data := fake.Course()

		sut := func() (*model.Course, error) {
			return model.NewCourse(data.ID, data.Name, data.CollegeID, data.College)
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

	s.Run("should return an error when 'collegeID' is zero", func() {
		sut := makeSut()

		sut.Data.CollegeID = 0

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
