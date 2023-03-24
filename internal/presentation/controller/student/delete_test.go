package controller_test

import (
	"fmt"
	"net/http"
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

type DeleteStudentSuite struct {
	suite.Suite
}

func TestDeleteStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteStudentSuite))
}

func (s *DeleteStudentSuite) TestHandle() {
	type Sut struct {
		Sut           controller.DeleteStudent
		Input         *service.DeleteInput
		DeleteStudent *mocks.DeleteStudent
	}

	makeSut := func() *Sut {
		input := fake.DeleteStudentInput()
		deleteStudent := new(mocks.DeleteStudent)
		sut := controller.NewDeleteStudent(deleteStudent)
		return &Sut{Sut: sut, DeleteStudent: deleteStudent, Input: input}
	}

	s.Run("should find one student", func() {
		sut := makeSut()

		sut.DeleteStudent.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteStudent.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteStudent.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteStudent.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
