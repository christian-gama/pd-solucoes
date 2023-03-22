package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/teacher"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllTeacherSuite struct {
	suite.Suite
}

func TestFindAllTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllTeacherSuite))
}

func (s *FindAllTeacherSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindAllTeachers
		TeacherRepo *mocks.Teacher
		Input       *service.FindAllTeachersInput
		Pagination  *querying.PaginationOutput[*model.Teacher]
	}

	makeSut := func() *Sut {
		teacherRepo := mocks.NewTeacher(s.T())
		sut := service.NewFindAllTeachers(teacherRepo)

		return &Sut{
			Sut:         sut,
			TeacherRepo: teacherRepo,
			Input:       fake.FindAllTeachersInput(),
			Pagination: &querying.PaginationOutput[*model.Teacher]{
				Total:   100,
				Results: []*model.Teacher{fakeModel.Teacher()},
			},
		}
	}

	s.Run("should find one teacher", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindAll", mock.Anything, mock.Anything).Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("teacherRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
