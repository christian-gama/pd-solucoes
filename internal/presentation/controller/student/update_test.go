package controller_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/student"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/student"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/student"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateStudentSuite struct {
	suite.Suite
}

func TestUpdateStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateStudentSuite))
}

func (s *UpdateStudentSuite) TestHandle() {
	type Sut struct {
		Sut           controller.UpdateStudent
		Input         *service.UpdateStudentInput
		UpdateStudent *mocks.UpdateStudent
	}

	makeSut := func() *Sut {
		input := fake.UpdateStudentInput()
		updateStudent := new(mocks.UpdateStudent)
		sut := controller.NewUpdateStudent(updateStudent)
		return &Sut{Sut: sut, UpdateStudent: updateStudent, Input: input}
	}

	s.Run("should update a student", func() {
		sut := makeSut()

		sut.UpdateStudent.On("Handle", mock.Anything, sut.Input).
			Return(&service.UpdateStudentOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateStudent.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when UpdateStudent.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateStudent.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
