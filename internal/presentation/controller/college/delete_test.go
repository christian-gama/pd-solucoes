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

type DeleteCollegeSuite struct {
	suite.Suite
}

func TestDeleteCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCollegeSuite))
}

func (s *DeleteCollegeSuite) TestHandle() {
	type Sut struct {
		Sut           controller.DeleteCollege
		Input         *service.DeleteInput
		DeleteCollege *mocks.DeleteCollege
	}

	makeSut := func() *Sut {
		input := fake.DeleteCollegeInput()
		deleteCollege := new(mocks.DeleteCollege)
		sut := controller.NewDeleteCollege(deleteCollege)
		return &Sut{Sut: sut, DeleteCollege: deleteCollege, Input: input}
	}

	s.Run("should find one college", func() {
		sut := makeSut()

		sut.DeleteCollege.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteCollege.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteCollege.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteCollege.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
