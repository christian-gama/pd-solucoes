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

type UpdateCourseSuite struct {
	suite.Suite
}

func TestUpdateCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateCourseSuite))
}

func (s *UpdateCourseSuite) TestHandle() {
	type Sut struct {
		Sut        service.UpdateCourse
		CourseRepo *mocks.Course
		Input      *service.UpdateCourseInput
		Course     *model.Course
	}

	makeSut := func() *Sut {
		courseRepo := mocks.NewCourse(s.T())
		sut := service.NewUpdateCourse(courseRepo)

		return &Sut{
			Sut:        sut,
			CourseRepo: courseRepo,
			Input:      fake.UpdateCourseInput(),
			Course:     fakeModel.Course(),
		}
	}

	s.Run("should add update a course", func() {
		sut := makeSut()

		sut.CourseRepo.On("Update", mock.Anything, mock.Anything).Return(sut.Course, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Course.ID, result.ID)
		s.Equal(sut.Course.Name, result.Name)
		s.Equal(sut.Course.CollegeID, result.CollegeID)
	})

	s.Run("courseRepo.Update returns an error", func() {
		sut := makeSut()

		sut.CourseRepo.On("Update", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})

	s.Run("should return an error if domain validation fails", func() {
		sut := makeSut()

		sut.Input.Name = ""

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.Error(err)
		s.Nil(result)
	})
}
