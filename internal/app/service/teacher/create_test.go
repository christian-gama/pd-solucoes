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

type CreateTeacherSuite struct {
	suite.Suite
}

func TestCreateTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateTeacherSuite))
}

func (s *CreateTeacherSuite) TestHandle() {
	type Sut struct {
		Sut         service.CreateTeacher
		TeacherRepo *mocks.Teacher
		Input       *service.CreateTeacherInput
		Teacher     *model.Teacher
	}

	makeSut := func() *Sut {
		teacherRepo := mocks.NewTeacher(s.T())
		sut := service.NewCreateTeacher(teacherRepo)

		return &Sut{
			Sut:         sut,
			TeacherRepo: teacherRepo,
			Input:       fake.CreateTeacherInput(),
			Teacher:     fakeModel.Teacher(),
		}
	}

	s.Run("should add a new teacher", func() {
		sut := makeSut()

		sut.TeacherRepo.On("Create", mock.Anything, mock.Anything).Return(sut.Teacher, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Teacher.ID, result.ID)
		s.Equal(sut.Teacher.Name, result.Name)
		s.Equal(sut.Teacher.Degree, result.Degree)
	})

	s.Run("teacherRepo.Create returns an error", func() {
		sut := makeSut()

		sut.TeacherRepo.On("Create", mock.Anything, mock.Anything).Return(nil, assert.AnError)

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
