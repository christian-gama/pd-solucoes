package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/courseSubject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneCourseSubjectSuite struct {
	suite.Suite
}

func TestFindOneCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCourseSubjectSuite))
}

func (s *FindOneCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut                  controller.FindOneCourseSubject
		Input                *service.FindOneInput
		FindOneCourseSubject *mocks.FindOneCourseSubject
	}

	makeSut := func() *Sut {
		input := fake.FindOneCourseSubjectInput()
		findOneCourseSubject := new(mocks.FindOneCourseSubject)
		sut := controller.NewFindOneCourseSubject(findOneCourseSubject)
		return &Sut{Sut: sut, FindOneCourseSubject: findOneCourseSubject, Input: input}
	}

	s.Run("should find one courseSubject", func() {
		sut := makeSut()

		sut.FindOneCourseSubject.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.Output{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneCourseSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneCourseSubject.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneCourseSubject.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
