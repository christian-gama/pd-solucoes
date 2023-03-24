package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/course"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/course"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/course"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCourseSuite struct {
	suite.Suite
}

func TestDeleteCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCourseSuite))
}

func (s *DeleteCourseSuite) TestHandle() {
	type Sut struct {
		Sut          controller.DeleteCourse
		Input        *service.DeleteInput
		DeleteCourse *mocks.DeleteCourse
	}

	makeSut := func() *Sut {
		input := fake.DeleteCourseInput()
		deleteCourse := new(mocks.DeleteCourse)
		sut := controller.NewDeleteCourse(deleteCourse)
		return &Sut{Sut: sut, DeleteCourse: deleteCourse, Input: input}
	}

	s.Run("should find one course", func() {
		sut := makeSut()

		sut.DeleteCourse.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteCourse.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteCourse.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteCourse.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
