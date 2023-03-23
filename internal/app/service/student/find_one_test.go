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

type FindOneStudentSuite struct {
	suite.Suite
}

func TestFindOneStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneStudentSuite))
}

func (s *FindOneStudentSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindOneStudent
		StudentRepo *mocks.Student
		Input       *service.FindOneStudentInput
		Student     *model.Student
	}

	makeSut := func() *Sut {
		studentRepo := mocks.NewStudent(s.T())
		sut := service.NewFindOneStudent(studentRepo)

		return &Sut{
			Sut:         sut,
			StudentRepo: studentRepo,
			Input:       fake.FindOneStudentInput(),
			Student:     fakeModel.Student(),
		}
	}

	s.Run("should find one student", func() {
		sut := makeSut()

		sut.StudentRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"courseSubjects",
				"courseSubjects.subject",
				"courseSubjects.course",
			).
			Return(sut.Student, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Student.ID, result.ID)
		s.Equal(sut.Student.Name, result.Name)
		s.Equal(sut.Student.Cpf, result.Cpf)
	})

	s.Run("studentRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.StudentRepo.
			On("FindOne", mock.Anything, mock.Anything,
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
