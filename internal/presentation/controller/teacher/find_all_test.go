package controller_test

import (
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/teacher"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllTeacherSuite struct {
	suite.Suite
}

func TestFindAllTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllTeacherSuite))
}

func (s *FindAllTeacherSuite) TestHandle() {
	type Sut struct {
		Sut             controller.FindAllTeachers
		Input           *service.FindAllInput
		FindAllTeachers *mocks.FindAllTeachers
	}

	makeSut := func() *Sut {
		input := fake.FindAllTeachersInput()
		findAllTeacher := new(mocks.FindAllTeachers)
		sut := controller.NewFindAllTeachers(findAllTeacher)
		return &Sut{Sut: sut, FindAllTeachers: findAllTeacher, Input: input}
	}

	s.Run("should find all teachers", func() {
		sut := makeSut()

		sut.FindAllTeachers.
			On("Handle", mock.Anything, mock.Anything).
			Return(&querying.PaginationOutput[*service.Output]{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllTeachers.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllTeacher.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllTeachers.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
