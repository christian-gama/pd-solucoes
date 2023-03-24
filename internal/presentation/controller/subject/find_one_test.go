package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/subject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/subject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/subject"
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
		Sut            controller.FindOneSubject
		Input          *service.FindOneInput
		FindOneSubject *mocks.FindOneSubject
	}

	makeSut := func() *Sut {
		input := fake.FindOneSubjectInput()
		findOneSubject := new(mocks.FindOneSubject)
		sut := controller.NewFindOneSubject(findOneSubject)
		return &Sut{Sut: sut, FindOneSubject: findOneSubject, Input: input}
	}

	s.Run("should find one subject", func() {
		sut := makeSut()

		sut.FindOneSubject.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.Output{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneSubject.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneSubject.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
