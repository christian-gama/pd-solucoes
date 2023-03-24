package controller_test

import (
	"net/http"
	"strings"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/college"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/college"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/college"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CreateCollegeSuite struct {
	suite.Suite
}

func TestCreateCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCollegeSuite))
}

func (s *CreateCollegeSuite) TestHandle() {
	type Sut struct {
		Sut           controller.CreateCollege
		Input         *service.CreateInput
		CreateCollege *mocks.CreateCollege
	}

	makeSut := func() *Sut {
		input := fake.CreateCollegeInput()
		createCollege := new(mocks.CreateCollege)
		sut := controller.NewCreateCollege(createCollege)
		return &Sut{Sut: sut, CreateCollege: createCollege, Input: input}
	}

	s.Run("should create a college", func() {
		sut := makeSut()

		sut.CreateCollege.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.CreateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateCollege.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when CreateCollege.Handle returns error", func() {
		sut := makeSut()

		sut.CreateCollege.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
