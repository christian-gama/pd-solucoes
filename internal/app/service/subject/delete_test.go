package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/subject"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteSubjectSuite struct {
	suite.Suite
}

func TestDeleteSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteSubjectSuite))
}

func (s *DeleteSubjectSuite) TestHandle() {
	type Sut struct {
		Sut         service.DeleteSubject
		SubjectRepo *mocks.Subject
		Input       *service.DeleteSubjectInput
		Subject     *model.Subject
	}

	makeSut := func() *Sut {
		subjectRepo := mocks.NewSubject(s.T())
		sut := service.NewDeleteSubject(subjectRepo)

		return &Sut{
			Sut:         sut,
			SubjectRepo: subjectRepo,
			Input:       fake.DeleteSubjectInput(),
		}
	}

	s.Run("should find one subject", func() {
		sut := makeSut()

		sut.SubjectRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Subject, nil)
		sut.SubjectRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("subjectRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.SubjectRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.Subject, nil)
		sut.SubjectRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("subjectRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.SubjectRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
