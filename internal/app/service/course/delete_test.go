package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/course"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCourseSuite struct {
	suite.Suite
}

func TestDeleteCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCourseSuite))
}

func (s *DeleteCourseSuite) TestHandle() {
	type Sut struct {
		Sut        service.DeleteCourse
		CourseRepo *mocks.Course
		Input      *service.DeleteCourseInput
		Course     *model.Course
	}

	makeSut := func() *Sut {
		courseRepo := mocks.NewCourse(s.T())
		sut := service.NewDeleteCourse(courseRepo)

		return &Sut{
			Sut:        sut,
			CourseRepo: courseRepo,
			Input:      fake.DeleteCourseInput(),
		}
	}

	s.Run("should find one course", func() {
		sut := makeSut()

		sut.CourseRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Course, nil)
		sut.CourseRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("courseRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.CourseRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Course, nil)
		sut.CourseRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("courseRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CourseRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
