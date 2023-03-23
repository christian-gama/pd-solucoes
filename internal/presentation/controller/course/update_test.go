package controller_test

import (
	"fmt"
	"net/http"
	"strings"
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

type UpdateCourseSuite struct {
	suite.Suite
}

func TestUpdateCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateCourseSuite))
}

func (s *UpdateCourseSuite) TestHandle() {
	type Sut struct {
		Sut          controller.UpdateCourse
		Input        *service.UpdateCourseInput
		UpdateCourse *mocks.UpdateCourse
	}

	makeSut := func() *Sut {
		input := fake.UpdateCourseInput()
		updateCourse := new(mocks.UpdateCourse)
		sut := controller.NewUpdateCourse(updateCourse)
		return &Sut{Sut: sut, UpdateCourse: updateCourse, Input: input}
	}

	s.Run("should update a course", func() {
		sut := makeSut()

		sut.UpdateCourse.On("Handle", mock.Anything, sut.Input).
			Return(&service.UpdateCourseOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateCourse.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when UpdateCourse.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateCourse.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
