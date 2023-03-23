package controller_test

import (
	"net/http"
	"strings"
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

type CreateTeacherSuite struct {
	suite.Suite
}

func TestCreateTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateTeacherSuite))
}

func (s *CreateTeacherSuite) TestHandle() {
	type Sut struct {
		Sut           controller.CreateTeacher
		Input         *service.CreateTeacherInput
		CreateTeacher *mocks.CreateTeacher
	}

	makeSut := func() *Sut {
		input := fake.CreateTeacherInput()
		createTeacher := new(mocks.CreateTeacher)
		sut := controller.NewCreateTeacher(createTeacher)
		return &Sut{Sut: sut, CreateTeacher: createTeacher, Input: input}
	}

	s.Run("should create a teacher", func() {
		sut := makeSut()

		sut.CreateTeacher.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.Output{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateTeacher.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when CreateTeacher.Handle returns error", func() {
		sut := makeSut()

		sut.CreateTeacher.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
