package controller_test

import (
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/courseSubject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCourseSubjectSuite struct {
	suite.Suite
}

func TestFindAllCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCourseSubjectSuite))
}

func (s *FindAllCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut                   controller.FindAllCourseSubjects
		Input                 *service.FindAllInput
		FindAllCourseSubjects *mocks.FindAllCourseSubjects
	}

	makeSut := func() *Sut {
		input := fake.FindAllCourseSubjectsInput()
		findAllCourseSubject := new(mocks.FindAllCourseSubjects)
		sut := controller.NewFindAllCourseSubjects(findAllCourseSubject)
		return &Sut{Sut: sut, FindAllCourseSubjects: findAllCourseSubject, Input: input}
	}

	s.Run("should find all courseSubjects", func() {
		sut := makeSut()

		sut.FindAllCourseSubjects.
			On("Handle", mock.Anything, mock.Anything).
			Return(&querying.PaginationOutput[*service.Output]{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllCourseSubjects.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllCourseSubject.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllCourseSubjects.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
