package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCourseSubjectSuite struct {
	suite.Suite
}

func TestDeleteCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCourseSubjectSuite))
}

func (s *DeleteCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut               service.DeleteCourseSubject
		CourseSubjectRepo *mocks.CourseSubject
		Input             *service.DeleleInput
		CourseSubject     *model.CourseSubject
	}

	makeSut := func() *Sut {
		courseSubjectRepo := mocks.NewCourseSubject(s.T())
		sut := service.NewDeleteCourseSubject(courseSubjectRepo)

		return &Sut{
			Sut:               sut,
			CourseSubjectRepo: courseSubjectRepo,
			Input:             fake.DeleteCourseSubjectInput(),
		}
	}

	s.Run("should find one courseSubject", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.On("FindOne", mock.Anything, mock.Anything).
			Return(sut.CourseSubject, nil)
		sut.CourseSubjectRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("courseSubjectRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.On("FindOne", mock.Anything, mock.Anything).
			Return(sut.CourseSubject, nil)
		sut.CourseSubjectRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("courseSubjectRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.On("FindOne", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
