package controller_test

import (
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/subject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/subject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/subject"
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
		Sut             controller.FindAllSubjects
		Input           *service.FindAllInput
		FindAllSubjects *mocks.FindAllSubjects
	}

	makeSut := func() *Sut {
		input := fake.FindAllSubjectsInput()
		findAllSubject := new(mocks.FindAllSubjects)
		sut := controller.NewFindAllSubjects(findAllSubject)
		return &Sut{Sut: sut, FindAllSubjects: findAllSubject, Input: input}
	}

	s.Run("should find all subjects", func() {
		sut := makeSut()

		sut.FindAllSubjects.
			On("Handle", mock.Anything, mock.Anything).
			Return(&querying.PaginationOutput[*service.Output]{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllSubjects.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllSubject.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllSubjects.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
