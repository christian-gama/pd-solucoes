package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/course"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneCourseSuite struct {
	suite.Suite
}

func TestFindOneCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCourseSuite))
}

func (s *FindOneCourseSuite) TestHandle() {
	type Sut struct {
		Sut        service.FindOneCourse
		CourseRepo *mocks.Course
		Input      *service.FindOneCourseInput
		Course     *model.Course
	}

	makeSut := func() *Sut {
		courseRepo := mocks.NewCourse(s.T())
		sut := service.NewFindOneCourse(courseRepo)

		return &Sut{
			Sut:        sut,
			CourseRepo: courseRepo,
			Input:      fake.FindOneCourseInput(),
			Course:     fakeModel.Course(),
		}
	}

	s.Run("should find one course", func() {
		sut := makeSut()

		sut.CourseRepo.
			On("FindOne", mock.Anything, mock.Anything, "enrollments", "subjects.students").
			Return(sut.Course, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Course.ID, result.ID)
		s.Equal(sut.Course.Name, result.Name)
	})

	s.Run("courseRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CourseRepo.
			On("FindOne", mock.Anything, mock.Anything, "enrollments", "subjects.students").
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
