package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/enrollment"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CreateCourseEnrollmentSuite struct {
	suite.Suite
}

func TestCreateCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCourseEnrollmentSuite))
}

func (s *CreateCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                  service.CreateCourseEnrollment
		CourseEnrollmentRepo *mocks.CourseEnrollment
		Input                *service.CreateInput
		CourseEnrollment     *model.CourseEnrollment
	}

	makeSut := func() *Sut {
		courseEnrollmentRepo := mocks.NewCourseEnrollment(s.T())
		sut := service.NewCreateCourseEnrollment(courseEnrollmentRepo)

		return &Sut{
			Sut:                  sut,
			CourseEnrollmentRepo: courseEnrollmentRepo,
			Input:                fake.CreateCourseEnrollmentInput(),
			CourseEnrollment:     fakeModel.CourseEnrollment(),
		}
	}

	s.Run("should add a new courseEnrollment", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.
			On("Create", mock.Anything, mock.Anything).
			Return(sut.CourseEnrollment, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.CourseEnrollment.ID, result.ID)
		s.Equal(sut.CourseEnrollment.CourseSubjectID, result.CourseSubjectID)
		s.Equal(sut.CourseEnrollment.StudentID, result.StudentID)
	})

	s.Run("courseEnrollmentRepo.Create returns an error", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.
			On("Create", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})

	s.Run("should return an error if domain validation fails", func() {
		sut := makeSut()

		sut.Input.CourseSubjectID = 0

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.Error(err)
		s.Nil(result)
	})
}
