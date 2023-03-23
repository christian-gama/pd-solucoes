package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/subject"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllSubjectSuite struct {
	suite.Suite
}

func TestFindAllSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllSubjectSuite))
}

func (s *FindAllSubjectSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindAllSubjects
		SubjectRepo *mocks.Subject
		Input       *service.FindAllSubjectsInput
		Pagination  *querying.PaginationOutput[*model.Subject]
	}

	makeSut := func() *Sut {
		subjectRepo := mocks.NewSubject(s.T())
		sut := service.NewFindAllSubjects(subjectRepo)

		return &Sut{
			Sut:         sut,
			SubjectRepo: subjectRepo,
			Input:       fake.FindAllSubjectsInput(),
			Pagination: &querying.PaginationOutput[*model.Subject]{
				Total:   100,
				Results: []*model.Subject{fakeModel.Subject()},
			},
		}
	}

	s.Run("should find all subjects", func() {
		sut := makeSut()

		sut.SubjectRepo.
			On("FindAll", mock.Anything, mock.Anything,
				"teacher",
				"courses",
				"courses.students",
			).
			Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("subjectRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.SubjectRepo.
			On("FindAll", mock.Anything, mock.Anything,
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
