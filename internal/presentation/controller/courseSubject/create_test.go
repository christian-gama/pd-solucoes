package controller_test

import (
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

type CreateCourseSubjectSuite struct {
	suite.Suite
}

func TestCreateCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCourseSubjectSuite))
}

func (s *CreateCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut                 controller.CreateCourseSubject
		Input               *service.CreateInput
		CreateCourseSubject *mocks.CreateCourseSubject
	}

	makeSut := func() *Sut {
		input := fake.CreateCourseSubjectInput()
		createCourseSubject := new(mocks.CreateCourseSubject)
		sut := controller.NewCreateCourseSubject(createCourseSubject)
		return &Sut{Sut: sut, CreateCourseSubject: createCourseSubject, Input: input}
	}

	s.Run("should create a course subject", func() {
		sut := makeSut()

		sut.CreateCourseSubject.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.CreateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateCourseSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid CourseID: it's required", func() {
		sut := makeSut()

		sut.Input.CourseID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid SubjectID: it's required", func() {
		sut := makeSut()

		sut.Input.SubjectID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when CreateCourseSubject.Handle returns error", func() {
		sut := makeSut()

		sut.CreateCourseSubject.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
