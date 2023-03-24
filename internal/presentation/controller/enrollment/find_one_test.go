package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/enrollment"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneCourseEnrollmentSuite struct {
	suite.Suite
}

func TestFindOneCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCourseEnrollmentSuite))
}

func (s *FindOneCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                     controller.FindOneCourseEnrollment
		Input                   *service.FindOneCourseEnrollmentInput
		FindOneCourseEnrollment *mocks.FindOneCourseEnrollment
	}

	makeSut := func() *Sut {
		input := fake.FindOneCourseEnrollmentInput()
		findOneCourseEnrollment := new(mocks.FindOneCourseEnrollment)
		sut := controller.NewFindOneCourseEnrollment(findOneCourseEnrollment)
		return &Sut{Sut: sut, FindOneCourseEnrollment: findOneCourseEnrollment, Input: input}
	}

	s.Run("should find one courseEnrollment", func() {
		sut := makeSut()

		sut.FindOneCourseEnrollment.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.Output{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindOneCourseEnrollment.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when FindOneCourseEnrollment.Handle returns error", func() {
		sut := makeSut()

		sut.FindOneCourseEnrollment.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
