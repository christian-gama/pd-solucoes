package controller_test

import (
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/enrollment"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCourseEnrollmentSuite struct {
	suite.Suite
}

func TestFindAllCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCourseEnrollmentSuite))
}

func (s *FindAllCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                      controller.FindAllCourseEnrollments
		Input                    *service.FindAllCourseEnrollmentsInput
		FindAllCourseEnrollments *mocks.FindAllCourseEnrollments
	}

	makeSut := func() *Sut {
		input := fake.FindAllCourseEnrollmentsInput()
		findAllCourseEnrollment := new(mocks.FindAllCourseEnrollments)
		sut := controller.NewFindAllCourseEnrollments(findAllCourseEnrollment)
		return &Sut{Sut: sut, FindAllCourseEnrollments: findAllCourseEnrollment, Input: input}
	}

	s.Run("should find all courseEnrollments", func() {
		sut := makeSut()

		sut.FindAllCourseEnrollments.
			On("Handle", mock.Anything, mock.Anything).
			Return(&querying.PaginationOutput[*service.Output]{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllCourseEnrollments.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllCourseEnrollment.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllCourseEnrollments.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
