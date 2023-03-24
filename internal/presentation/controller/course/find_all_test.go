package controller_test

import (
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/course"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/course"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/course"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/course"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCourseSuite struct {
	suite.Suite
}

func TestFindAllCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCourseSuite))
}

func (s *FindAllCourseSuite) TestHandle() {
	type Sut struct {
		Sut            controller.FindAllCourses
		Input          *service.FindAllInput
		FindAllCourses *mocks.FindAllCourses
	}

	makeSut := func() *Sut {
		input := fake.FindAllCoursesInput()
		findAllCourse := new(mocks.FindAllCourses)
		sut := controller.NewFindAllCourses(findAllCourse)
		return &Sut{Sut: sut, FindAllCourses: findAllCourse, Input: input}
	}

	s.Run("should find all courses", func() {
		sut := makeSut()

		sut.FindAllCourses.
			On("Handle", mock.Anything, mock.Anything).
			Return(&querying.PaginationOutput[*service.Output]{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllCourses.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllCourse.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllCourses.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
