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

type FindOneCourseEnrollmentSuite struct {
	suite.Suite
}

func TestFindOneCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCourseEnrollmentSuite))
}

func (s *FindOneCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                  service.FindOneCourseEnrollment
		CourseEnrollmentRepo *mocks.CourseEnrollment
		Input                *service.FindOneInput
		CourseEnrollment     *model.CourseEnrollment
	}

	makeSut := func() *Sut {
		courseEnrollmentRepo := mocks.NewCourseEnrollment(s.T())
		sut := service.NewFindOneCourseEnrollment(courseEnrollmentRepo)

		return &Sut{
			Sut:                  sut,
			CourseEnrollmentRepo: courseEnrollmentRepo,
			Input:                fake.FindOneCourseEnrollmentInput(),
			CourseEnrollment:     fakeModel.CourseEnrollment(),
		}
	}

	s.Run("should find one courseEnrollment", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"courseSubject",
				"courseSubject.subject",
				"courseSubject.course",
			).
			Return(sut.CourseEnrollment, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.CourseEnrollment.ID, result.ID)
		s.Equal(sut.CourseEnrollment.CourseSubjectID, result.CourseSubject.ID)
	})

	s.Run("courseEnrollmentRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"courseSubject",
				"courseSubject.subject",
				"courseSubject.course",
			).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
