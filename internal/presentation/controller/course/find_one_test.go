package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/course"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/course"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/course"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneCourseSuite struct {
	suite.Suite
}

func TestFindOneCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCourseSuite))
}

func (s *FindOneCourseSuite) TestHandle() {
	type Sut struct {
		Sut           controller.FindOneCourse
		Input         *service.FindOneCourseInput
		FindOneCourse *mocks.FindOneCourse
	}

	makeSut := func() *Sut {
		input := fake.FindOneCourseInput()
		findOneCourse := new(mocks.FindOneCourse)
		sut := controller.NewFindOneCourse(findOneCourse)
		return &Sut{Sut: sut, FindOneCourse: findOneCourse, Input: input}
	}

	s.Run("should find one course", func() {
		sut := makeSut()

		sut.FindOneCourse.
			On("Handle", mock.Anything, sut.Input).
			Return(&model.Course{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneCourse.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneCourse.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneCourse.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
