package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/teacher"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteTeacherSuite struct {
	suite.Suite
}

func TestDeleteTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteTeacherSuite))
}

func (s *DeleteTeacherSuite) TestHandle() {
	type Sut struct {
		Sut         service.DeleteTeacher
		TeacherRepo *mocks.Teacher
		Input       *service.DeleteTeacherInput
		Teacher     *model.Teacher
	}

	makeSut := func() *Sut {
		teacherRepo := mocks.NewTeacher(s.T())
		sut := service.NewDeleteTeacher(teacherRepo)

		return &Sut{
			Sut:         sut,
			TeacherRepo: teacherRepo,
			Input:       fake.DeleteTeacherInput(),
		}
	}

	s.Run("should find one teacher", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Teacher, nil)
		sut.TeacherRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("teacherRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Teacher, nil)
		sut.TeacherRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("teacherRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.TeacherRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
