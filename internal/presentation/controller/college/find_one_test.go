package controller_test

import (
	"fmt"
	"net/http"
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

type FindOneCollegeSuite struct {
	suite.Suite
}

func TestFindOneCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCollegeSuite))
}

func (s *FindOneCollegeSuite) TestHandle() {
	type Sut struct {
		Sut            controller.FindOneCollege
		Input          *service.FindOneInput
		FindOneCollege *mocks.FindOneCollege
	}

	makeSut := func() *Sut {
		input := fake.FindOneCollegeInput()
		findOneCollege := new(mocks.FindOneCollege)
		sut := controller.NewFindOneCollege(findOneCollege)
		return &Sut{Sut: sut, FindOneCollege: findOneCollege, Input: input}
	}

	s.Run("should find one college", func() {
		sut := makeSut()

		sut.FindOneCollege.On("Handle", mock.Anything, sut.Input).
			Return(&service.Output{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneCollege.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneCollege.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneCollege.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
