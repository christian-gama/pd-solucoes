package controller_test

import (
	"net/http"
	"strings"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/student"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/student"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/student"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CreateStudentSuite struct {
	suite.Suite
}

func TestCreateStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateStudentSuite))
}

func (s *CreateStudentSuite) TestHandle() {
	type Sut struct {
		Sut           controller.CreateStudent
		Input         *service.CreateStudentInput
		CreateStudent *mocks.CreateStudent
	}

	makeSut := func() *Sut {
		input := fake.CreateStudentInput()
		createStudent := new(mocks.CreateStudent)
		sut := controller.NewCreateStudent(createStudent)
		return &Sut{Sut: sut, CreateStudent: createStudent, Input: input}
	}

	s.Run("should create a student", func() {
		sut := makeSut()

		sut.CreateStudent.
			On("Handle", mock.Anything, sut.Input).
			Return(&model.Student{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateStudent.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when CreateStudent.Handle returns error", func() {
		sut := makeSut()

		sut.CreateStudent.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
