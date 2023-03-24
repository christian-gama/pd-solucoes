package controller_test

import (
	"fmt"
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

type UpdateSubjectSuite struct {
	suite.Suite
}

func TestUpdateSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateSubjectSuite))
}

func (s *UpdateSubjectSuite) TestHandle() {
	type Sut struct {
		Sut           controller.UpdateSubject
		Input         *service.UpdateInput
		UpdateSubject *mocks.UpdateSubject
	}

	makeSut := func() *Sut {
		input := fake.UpdateSubjectInput()
		updateSubject := new(mocks.UpdateSubject)
		sut := controller.NewUpdateSubject(updateSubject)
		return &Sut{Sut: sut, UpdateSubject: updateSubject, Input: input}
	}

	s.Run("should update a subject", func() {
		sut := makeSut()

		sut.UpdateSubject.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.UpdateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when UpdateSubject.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateSubject.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
