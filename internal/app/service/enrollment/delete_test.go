package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/enrollment"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCourseEnrollmentSuite struct {
	suite.Suite
}

func TestDeleteCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCourseEnrollmentSuite))
}

func (s *DeleteCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                  service.DeleteCourseEnrollment
		CourseEnrollmentRepo *mocks.CourseEnrollment
		Input                *service.DeleteCourseEnrollmentInput
		CourseEnrollment     *model.CourseEnrollment
	}

	makeSut := func() *Sut {
		courseEnrollmentRepo := mocks.NewCourseEnrollment(s.T())
		sut := service.NewDeleteCourseEnrollment(courseEnrollmentRepo)

		return &Sut{
			Sut:                  sut,
			CourseEnrollmentRepo: courseEnrollmentRepo,
			Input:                fake.DeleteCourseEnrollmentInput(),
		}
	}

	s.Run("should find one courseEnrollment", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.On("FindOne", mock.Anything, mock.Anything).
			Return(sut.CourseEnrollment, nil)
		sut.CourseEnrollmentRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("courseEnrollmentRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.On("FindOne", mock.Anything, mock.Anything).
			Return(sut.CourseEnrollment, nil)
		sut.CourseEnrollmentRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("courseEnrollmentRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
