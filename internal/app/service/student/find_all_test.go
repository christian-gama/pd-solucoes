package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/student"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllStudentSuite struct {
	suite.Suite
}

func TestFindAllStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllStudentSuite))
}

func (s *FindAllStudentSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindAllStudents
		StudentRepo *mocks.Student
		Input       *service.FindAllStudentsInput
		Pagination  *querying.PaginationOutput[*model.Student]
	}

	makeSut := func() *Sut {
		studentRepo := mocks.NewStudent(s.T())
		sut := service.NewFindAllStudents(studentRepo)

		return &Sut{
			Sut:         sut,
			StudentRepo: studentRepo,
			Input:       fake.FindAllStudentsInput(),
			Pagination: &querying.PaginationOutput[*model.Student]{
				Total: 100,
				Results: []*model.Student{
					fakeModel.Student(),
				},
			},
		}
	}

	s.Run("should find all students", func() {
		sut := makeSut()

		sut.StudentRepo.
			On("FindAll", mock.Anything, mock.Anything,
				"courseSubjects",
				"courseSubjects.subject",
				"courseSubjects.course",
			).
			Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("studentRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.StudentRepo.
			On("FindAll", mock.Anything, mock.Anything,
				"courseSubjects",
				"courseSubjects.subject",
				"courseSubjects.course",
			).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
