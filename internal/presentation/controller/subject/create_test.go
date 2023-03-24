package controller_test

import (
	"net/http"
	"strings"
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

type CreateSubjectSuite struct {
	suite.Suite
}

func TestCreateSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateSubjectSuite))
}

func (s *CreateSubjectSuite) TestHandle() {
	type Sut struct {
		Sut           controller.CreateSubject
		Input         *service.CreateInput
		CreateSubject *mocks.CreateSubject
	}

	makeSut := func() *Sut {
		input := fake.CreateSubjectInput()
		createSubject := new(mocks.CreateSubject)
		sut := controller.NewCreateSubject(createSubject)
		return &Sut{Sut: sut, CreateSubject: createSubject, Input: input}
	}

	s.Run("should create a subject", func() {
		sut := makeSut()

		sut.CreateSubject.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.CreateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid Name: it's required", func() {
		sut := makeSut()

		sut.Input.Name = ""

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid Name: max length", func() {
		sut := makeSut()

		sut.Input.Name = strings.Repeat("a", 101)

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when CreateSubject.Handle returns error", func() {
		sut := makeSut()

		sut.CreateSubject.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
