package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/student"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateStudentSuite struct {
	suite.Suite
}

func TestUpdateStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateStudentSuite))
}

func (s *UpdateStudentSuite) TestHandle() {
	type Sut struct {
		Sut         service.UpdateStudent
		StudentRepo *mocks.Student
		Input       *service.UpdateInput
		Student     *model.Student
	}

	makeSut := func() *Sut {
		studentRepo := mocks.NewStudent(s.T())
		sut := service.NewUpdateStudent(studentRepo)

		return &Sut{
			Sut:         sut,
			StudentRepo: studentRepo,
			Input:       fake.UpdateStudentInput(),
			Student:     fakeModel.Student(),
		}
	}

	s.Run("should add update a student", func() {
		sut := makeSut()

		sut.StudentRepo.On("Update", mock.Anything, mock.Anything).Return(sut.Student, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Student.ID, result.ID)
		s.Equal(sut.Student.Name, result.Name)
		s.Equal(sut.Student.Cpf, result.Cpf)
	})

	s.Run("studentRepo.Update returns an error", func() {
		sut := makeSut()

		sut.StudentRepo.On("Update", mock.Anything, mock.Anything).Return(nil, assert.AnError)

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
