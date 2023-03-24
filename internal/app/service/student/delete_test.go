package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/student"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteStudentSuite struct {
	suite.Suite
}

func TestDeleteStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteStudentSuite))
}

func (s *DeleteStudentSuite) TestHandle() {
	type Sut struct {
		Sut         service.DeleteStudent
		StudentRepo *mocks.Student
		Input       *service.DeleteInput
		Student     *model.Student
	}

	makeSut := func() *Sut {
		studentRepo := mocks.NewStudent(s.T())
		sut := service.NewDeleteStudent(studentRepo)

		return &Sut{
			Sut:         sut,
			StudentRepo: studentRepo,
			Input:       fake.DeleteStudentInput(),
		}
	}

	s.Run("should find one student", func() {
		sut := makeSut()

		sut.StudentRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Student, nil)
		sut.StudentRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("studentRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.StudentRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Student, nil)
		sut.StudentRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("studentRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.StudentRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
