package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneCourseSubjectSuite struct {
	suite.Suite
}

func TestFindOneCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCourseSubjectSuite))
}

func (s *FindOneCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut               service.FindOneCourseSubject
		CourseSubjectRepo *mocks.CourseSubject
		Input             *service.FindOneCourseSubjectInput
		CourseSubject     *model.CourseSubject
	}

	makeSut := func() *Sut {
		courseSubjectRepo := mocks.NewCourseSubject(s.T())
		sut := service.NewFindOneCourseSubject(courseSubjectRepo)

		return &Sut{
			Sut:               sut,
			CourseSubjectRepo: courseSubjectRepo,
			Input:             fake.FindOneCourseSubjectInput(),
			CourseSubject:     fakeModel.CourseSubject(),
		}
	}

	s.Run("should find one courseSubject", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"subject",
				"course",
				"students",
			).
			Return(sut.CourseSubject, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.CourseSubject.ID, result.ID)
		s.Equal(sut.CourseSubject.CourseID, result.Course.ID)
	})

	s.Run("courseSubjectRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"subject",
				"course",
				"students",
			).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
