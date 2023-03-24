package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCourseSubjectSuite struct {
	suite.Suite
}

func TestFindAllCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCourseSubjectSuite))
}

func (s *FindAllCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut               service.FindAllCourseSubjects
		CourseSubjectRepo *mocks.CourseSubject
		Input             *service.FindAllInput
		Pagination        *querying.PaginationOutput[*model.CourseSubject]
	}

	makeSut := func() *Sut {
		courseSubjectRepo := mocks.NewCourseSubject(s.T())
		sut := service.NewFindAllCourseSubjects(courseSubjectRepo)

		return &Sut{
			Sut:               sut,
			CourseSubjectRepo: courseSubjectRepo,
			Input:             fake.FindAllCourseSubjectsInput(),
			Pagination: &querying.PaginationOutput[*model.CourseSubject]{
				Total:   100,
				Results: []*model.CourseSubject{fakeModel.CourseSubject()},
			},
		}
	}

	s.Run("should find all courseSubjects", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("FindAll", mock.Anything, mock.Anything,
				"subject",
				"course",
				"students",
			).
			Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("courseSubjectRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("FindAll", mock.Anything, mock.Anything,
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
