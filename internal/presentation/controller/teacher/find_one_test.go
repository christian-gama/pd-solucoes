package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/teacher"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/teacher"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneTeacherSuite struct {
	suite.Suite
}

func TestFindOneTeacherSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneTeacherSuite))
}

func (s *FindOneTeacherSuite) TestHandle() {
	type Sut struct {
		Sut            controller.FindOneTeacher
		Input          *service.FindOneTeacherInput
		FindOneTeacher *mocks.FindOneTeacher
	}

	makeSut := func() *Sut {
		input := fake.FindOneTeacherInput()
		findOneTeacher := new(mocks.FindOneTeacher)
		sut := controller.NewFindOneTeacher(findOneTeacher)
		return &Sut{Sut: sut, FindOneTeacher: findOneTeacher, Input: input}
	}

	s.Run("should find one teacher", func() {
		sut := makeSut()

		sut.FindOneTeacher.
			On("Handle", mock.Anything, sut.Input).
			Return(&model.Teacher{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneTeacher.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneTeacher.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneTeacher.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
