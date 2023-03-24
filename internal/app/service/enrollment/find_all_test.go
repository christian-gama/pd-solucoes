package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/enrollment"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCourseEnrollmentSuite struct {
	suite.Suite
}

func TestFindAllCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCourseEnrollmentSuite))
}

func (s *FindAllCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                  service.FindAllCourseEnrollments
		CourseEnrollmentRepo *mocks.CourseEnrollment
		Input                *service.FindAllInput
		Pagination           *querying.PaginationOutput[*model.CourseEnrollment]
	}

	makeSut := func() *Sut {
		courseEnrollmentRepo := mocks.NewCourseEnrollment(s.T())
		sut := service.NewFindAllCourseEnrollments(courseEnrollmentRepo)

		return &Sut{
			Sut:                  sut,
			CourseEnrollmentRepo: courseEnrollmentRepo,
			Input:                fake.FindAllCourseEnrollmentsInput(),
			Pagination: &querying.PaginationOutput[*model.CourseEnrollment]{
				Total:   100,
				Results: []*model.CourseEnrollment{fakeModel.CourseEnrollment()},
			},
		}
	}

	s.Run("should find all courseEnrollments", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.
			On("FindAll", mock.Anything, mock.Anything,
				"courseSubject",
				"courseSubject.subject",
				"courseSubject.course",
			).
			Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("courseEnrollmentRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.CourseEnrollmentRepo.
			On("FindAll", mock.Anything, mock.Anything,
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
