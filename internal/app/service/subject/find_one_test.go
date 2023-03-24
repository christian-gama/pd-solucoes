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

type FindOneSubjectSuite struct {
	suite.Suite
}

func TestFindOneSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneSubjectSuite))
}

func (s *FindOneSubjectSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindOneSubject
		SubjectRepo *mocks.Subject
		Input       *service.FindOneInput
		Subject     *model.Subject
	}

	makeSut := func() *Sut {
		subjectRepo := mocks.NewSubject(s.T())
		sut := service.NewFindOneSubject(subjectRepo)

		return &Sut{
			Sut:         sut,
			SubjectRepo: subjectRepo,
			Input:       fake.FindOneSubjectInput(),
			Subject:     fakeModel.Subject(),
		}
	}

	s.Run("should find one subject", func() {
		sut := makeSut()

		sut.SubjectRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"teacher",
				"courses",
				"courses.students",
			).
			Return(sut.Subject, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Subject.ID, result.ID)
		s.Equal(sut.Subject.Name, result.Name)
	})

	s.Run("subjectRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.SubjectRepo.
			On("FindOne", mock.Anything, mock.Anything,
				"teacher",
				"courses",
				"courses.students",
			).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
