package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/subject"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CreateSubjectSuite struct {
	suite.Suite
}

func TestCreateSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateSubjectSuite))
}

func (s *CreateSubjectSuite) TestHandle() {
	type Sut struct {
		Sut         service.CreateSubject
		SubjectRepo *mocks.Subject
		Input       *service.CreateSubjectInput
		Subject     *model.Subject
	}

	makeSut := func() *Sut {
		subjectRepo := mocks.NewSubject(s.T())
		sut := service.NewCreateSubject(subjectRepo)

		return &Sut{
			Sut:         sut,
			SubjectRepo: subjectRepo,
			Input:       fake.CreateSubjectInput(),
			Subject:     fakeModel.Subject(),
		}
	}

	s.Run("should add a new subject", func() {
		sut := makeSut()

		sut.SubjectRepo.On("Create", mock.Anything, mock.Anything).Return(sut.Subject, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Subject.ID, result.ID)
		s.Equal(sut.Subject.Name, result.Name)
		s.Equal(sut.Subject.TeacherID, result.Teacher.ID)
	})

	s.Run("subjectRepo.Create returns an error", func() {
		sut := makeSut()

		sut.SubjectRepo.On("Create", mock.Anything, mock.Anything).Return(nil, assert.AnError)

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
