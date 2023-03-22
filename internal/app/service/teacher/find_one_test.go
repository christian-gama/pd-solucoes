package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/teacher"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneTeacherSuite struct {
	suite.Suite
}

func TestFindOneTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneTeacherSuite))
}

func (s *FindOneTeacherSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindOneTeacher
		TeacherRepo *mocks.Teacher
		Input       *service.FindOneTeacherInput
		Teacher     *model.Teacher
	}

	makeSut := func() *Sut {
		teacherRepo := mocks.NewTeacher(s.T())
		sut := service.NewFindOneTeacher(teacherRepo)

		return &Sut{
			Sut:         sut,
			TeacherRepo: teacherRepo,
			Input:       fake.FindOneTeacherInput(),
			Teacher:     fakeModel.Teacher(),
		}
	}

	s.Run("should find one teacher", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Teacher, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Teacher.ID, result.ID)
		s.Equal(sut.Teacher.Name, result.Name)
		s.Equal(sut.Teacher.Degree, result.Degree)
	})

	s.Run("teacherRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
