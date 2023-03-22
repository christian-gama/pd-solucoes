package controller_test

import (
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/student"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/student"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/student"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllStudentSuite struct {
	suite.Suite
}

func TestFindAllStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllStudentSuite))
}

func (s *FindAllStudentSuite) TestHandle() {
	type Sut struct {
		Sut             controller.FindAllStudents
		Input           *service.FindAllStudentsInput
		FindAllStudents *mocks.FindAllStudents
	}

	makeSut := func() *Sut {
		input := fake.FindAllStudentsInput()
		findAllStudent := new(mocks.FindAllStudents)
		sut := controller.NewFindAllStudents(findAllStudent)
		return &Sut{Sut: sut, FindAllStudents: findAllStudent, Input: input}
	}

	s.Run("should find all students", func() {
		sut := makeSut()

		sut.FindAllStudents.On("Handle", mock.Anything, mock.Anything).
			Return(&service.FindAllStudentsOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllStudents.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllStudent.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllStudents.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
