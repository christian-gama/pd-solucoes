package controller_test

import (
	"fmt"
	"net/http"
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

type DeleteTeacherSuite struct {
	suite.Suite
}

func TestDeleteTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteTeacherSuite))
}

func (s *DeleteTeacherSuite) TestHandle() {
	type Sut struct {
		Sut           controller.DeleteTeacher
		Input         *service.DeleteTeacherInput
		DeleteTeacher *mocks.DeleteTeacher
	}

	makeSut := func() *Sut {
		input := fake.DeleteTeacherInput()
		deleteTeacher := new(mocks.DeleteTeacher)
		sut := controller.NewDeleteTeacher(deleteTeacher)
		return &Sut{Sut: sut, DeleteTeacher: deleteTeacher, Input: input}
	}

	s.Run("should find one teacher", func() {
		sut := makeSut()

		sut.DeleteTeacher.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteTeacher.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteTeacher.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteTeacher.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
