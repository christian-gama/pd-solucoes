package controller_test

import (
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

type CreateCourseSuite struct {
	suite.Suite
}

func TestCreateCourseSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCourseSuite))
}

func (s *CreateCourseSuite) TestHandle() {
	type Sut struct {
		Sut          controller.CreateCourse
		Input        *service.CreateCourseInput
		CreateCourse *mocks.CreateCourse
	}

	makeSut := func() *Sut {
		input := fake.CreateCourseInput()
		createCourse := new(mocks.CreateCourse)
		sut := controller.NewCreateCourse(createCourse)
		return &Sut{Sut: sut, CreateCourse: createCourse, Input: input}
	}

	s.Run("should create a course", func() {
		sut := makeSut()

		sut.CreateCourse.On("Handle", mock.Anything, sut.Input).
			Return(&service.CreateCourseOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateCourse.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid Name: it's required", func() {
		sut := makeSut()

		sut.Input.Name = ""

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid Name: max length", func() {
		sut := makeSut()

		sut.Input.Name = strings.Repeat("a", 101)

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when CreateCourse.Handle returns error", func() {
		sut := makeSut()

		sut.CreateCourse.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
