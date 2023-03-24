package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/courseSubject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCourseSubjectSuite struct {
	suite.Suite
}

func TestDeleteCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCourseSubjectSuite))
}

func (s *DeleteCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut                 controller.DeleteCourseSubject
		Input               *service.DeleleInput
		DeleteCourseSubject *mocks.DeleteCourseSubject
	}

	makeSut := func() *Sut {
		input := fake.DeleteCourseSubjectInput()
		deleteCourseSubject := new(mocks.DeleteCourseSubject)
		sut := controller.NewDeleteCourseSubject(deleteCourseSubject)
		return &Sut{Sut: sut, DeleteCourseSubject: deleteCourseSubject, Input: input}
	}

	s.Run("should find one courseSubject", func() {
		sut := makeSut()

		sut.DeleteCourseSubject.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteCourseSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteCourseSubject.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteCourseSubject.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
