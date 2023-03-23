package controller_test

import (
	"fmt"
	"net/http"
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

type FindOneStudentSuite struct {
	suite.Suite
}

func TestFindOneStudentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneStudentSuite))
}

func (s *FindOneStudentSuite) TestHandle() {
	type Sut struct {
		Sut            controller.FindOneStudent
		Input          *service.FindOneStudentInput
		FindOneStudent *mocks.FindOneStudent
	}

	makeSut := func() *Sut {
		input := fake.FindOneStudentInput()
		findOneStudent := new(mocks.FindOneStudent)
		sut := controller.NewFindOneStudent(findOneStudent)
		return &Sut{Sut: sut, FindOneStudent: findOneStudent, Input: input}
	}

	s.Run("should find one student", func() {
		sut := makeSut()

		sut.FindOneStudent.
			On("Handle", mock.Anything, sut.Input).
			Return(&model.Student{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneStudent.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneStudent.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneStudent.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
