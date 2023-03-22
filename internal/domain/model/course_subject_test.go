package model_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CourseSubjectSuite struct {
	suite.Suite
}

func TestCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CourseSubjectSuite))
}

func (s *CourseSubjectSuite) TestNewCourseSubject() {
	type Sut struct {
		Sut  func() (*model.CourseSubject, error)
		Data *model.CourseSubject
	}

	makeSut := func() *Sut {
		data := fake.CourseSubject()

		sut := func() (*model.CourseSubject, error) {
			return model.NewCourseSubject(
				data.ID,
				data.CourseID,
				data.SubjectID,
				data.Course,
				data.Subject,
			)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("should create a new model", func() {
		sut := makeSut()

		model, err := sut.Sut()

		s.NoError(err)
		s.NotNil(model, "model should not be nil")
	})

	s.Run("should return an error when 'courseID' is zero", func() {
		sut := makeSut()

		sut.Data.CourseID = 0

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'subjectID' is zero", func() {
		sut := makeSut()

		sut.Data.SubjectID = 0

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'course' is nil", func() {
		sut := makeSut()

		sut.Data.Course = nil

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'course' is invalid", func() {
		sut := makeSut()

		sut.Data.Course.Name = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'subject' is nil", func() {
		sut := makeSut()

		sut.Data.Subject = nil

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'subject' is invalid", func() {
		sut := makeSut()

		sut.Data.Subject.Name = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
