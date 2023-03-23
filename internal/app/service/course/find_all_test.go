package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/course"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCourseSuite struct {
	suite.Suite
}

func TestFindAllCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCourseSuite))
}

func (s *FindAllCourseSuite) TestHandle() {
	type Sut struct {
		Sut        service.FindAllCourses
		CourseRepo *mocks.Course
		Input      *service.FindAllCoursesInput
		Pagination *querying.PaginationOutput[*model.Course]
	}

	makeSut := func() *Sut {
		courseRepo := mocks.NewCourse(s.T())
		sut := service.NewFindAllCourses(courseRepo)

		return &Sut{
			Sut:        sut,
			CourseRepo: courseRepo,
			Input:      fake.FindAllCoursesInput(),
			Pagination: &querying.PaginationOutput[*model.Course]{
				Total:   100,
				Results: []*model.Course{fakeModel.Course()},
			},
		}
	}

	s.Run("should find one course", func() {
		sut := makeSut()

		sut.CourseRepo.
			On("FindAll", mock.Anything, mock.Anything, "enrollments", "subjects.students").
			Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("courseRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.CourseRepo.
			On("FindAll", mock.Anything, mock.Anything, "enrollments", "subjects.students").
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
