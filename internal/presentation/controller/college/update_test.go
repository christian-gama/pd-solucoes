package controller_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/college"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/college"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/college"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateCollegeSuite struct {
	suite.Suite
}

func TestUpdateCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateCollegeSuite))
}

func (s *UpdateCollegeSuite) TestHandle() {
	type Sut struct {
		Sut           controller.UpdateCollege
		Input         *service.UpdateCollegeInput
		UpdateCollege *mocks.UpdateCollege
	}

	makeSut := func() *Sut {
		input := fake.UpdateCollegeInput()
		updateCollege := new(mocks.UpdateCollege)
		sut := controller.NewUpdateCollege(updateCollege)
		return &Sut{Sut: sut, UpdateCollege: updateCollege, Input: input}
	}

	s.Run("should update a college", func() {
		sut := makeSut()

		sut.UpdateCollege.
			On("Handle", mock.Anything, sut.Input).
			Return(&model.College{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateCollege.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid Name: it's required", func() {
		sut := makeSut()

		sut.Input.Name = ""

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid Name: max length", func() {
		sut := makeSut()

		sut.Input.Name = strings.Repeat("a", 101)

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when UpdateCollege.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateCollege.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
