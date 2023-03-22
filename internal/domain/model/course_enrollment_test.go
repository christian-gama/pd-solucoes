package model_test

import (
	"testing"
	"time"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CourseEnrollmentSuite struct {
	suite.Suite
}

func TestCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CourseEnrollmentSuite))
}

func (s *CourseEnrollmentSuite) TestNewCourseEnrollment() {
	type Sut struct {
		Sut  func() (*model.CourseEnrollment, error)
		Data *model.CourseEnrollment
	}

	makeSut := func() *Sut {
		data := fake.CourseEnrollment()

		sut := func() (*model.CourseEnrollment, error) {
			return model.NewCourseEnrollment(
				data.ID,
				data.StudentID,
				data.Student,
				data.EnrollmentDate,
				data.CourseSubjectID,
				data.CourseSubject,
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

	s.Run("should return an error when 'studentID' is zero", func() {
		sut := makeSut()

		sut.Data.StudentID = 0

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'student' is nil", func() {
		sut := makeSut()

		sut.Data.Student = nil

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'student' is invalid", func() {
		sut := makeSut()

		sut.Data.Student.Name = ""

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'enrollmentDate' is zero", func() {
		sut := makeSut()

		sut.Data.EnrollmentDate = time.Time{}

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'enrollmentDate' is in the future", func() {
		sut := makeSut()

		sut.Data.EnrollmentDate = time.Now().Add(time.Hour * 24)

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'courseSubjectID' is zero", func() {
		sut := makeSut()

		sut.Data.CourseSubjectID = 0

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})

	s.Run("should return an error when 'courseSubject' is nil", func() {
		sut := makeSut()

		sut.Data.CourseSubject = nil

		model, err := sut.Sut()

		s.Error(err)
		s.Nil(model, "model should be nil")
	})
}
