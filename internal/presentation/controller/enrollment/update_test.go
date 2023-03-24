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

type UpdateCourseEnrollmentSuite struct {
	suite.Suite
}

func TestUpdateCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateCourseEnrollmentSuite))
}

func (s *UpdateCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                    controller.UpdateCourseEnrollment
		Input                  *service.UpdateInput
		UpdateCourseEnrollment *mocks.UpdateCourseEnrollment
	}

	makeSut := func() *Sut {
		input := fake.UpdateCourseEnrollmentInput()
		updateCourseEnrollment := new(mocks.UpdateCourseEnrollment)
		sut := controller.NewUpdateCourseEnrollment(updateCourseEnrollment)
		return &Sut{Sut: sut, UpdateCourseEnrollment: updateCourseEnrollment, Input: input}
	}

	s.Run("should update a courseEnrollment", func() {
		sut := makeSut()

		sut.UpdateCourseEnrollment.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.UpdateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateCourseEnrollment.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid StudentID: it's required", func() {
		sut := makeSut()

		sut.Input.StudentID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid CourseSubjectID: it's required", func() {
		sut := makeSut()

		sut.Input.CourseSubjectID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when UpdateCourseEnrollment.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateCourseEnrollment.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
