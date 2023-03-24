package controller_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/teacher"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateTeacherSuite struct {
	suite.Suite
}

func TestUpdateTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateTeacherSuite))
}

func (s *UpdateTeacherSuite) TestHandle() {
	type Sut struct {
		Sut           controller.UpdateTeacher
		Input         *service.UpdateInput
		UpdateTeacher *mocks.UpdateTeacher
	}

	makeSut := func() *Sut {
		input := fake.UpdateTeacherInput()
		updateTeacher := new(mocks.UpdateTeacher)
		sut := controller.NewUpdateTeacher(updateTeacher)
		return &Sut{Sut: sut, UpdateTeacher: updateTeacher, Input: input}
	}

	s.Run("should update a teacher", func() {
		sut := makeSut()

		sut.UpdateTeacher.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.UpdateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateTeacher.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when UpdateTeacher.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateTeacher.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
