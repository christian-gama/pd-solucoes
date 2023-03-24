package controller_test

import (
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

type CreateCourseEnrollmentSuite struct {
	suite.Suite
}

func TestCreateCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCourseEnrollmentSuite))
}

func (s *CreateCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                    controller.CreateCourseEnrollment
		Input                  *service.CreateCourseEnrollmentInput
		CreateCourseEnrollment *mocks.CreateCourseEnrollment
	}

	makeSut := func() *Sut {
		input := fake.CreateCourseEnrollmentInput()
		createCourseEnrollment := new(mocks.CreateCourseEnrollment)
		sut := controller.NewCreateCourseEnrollment(createCourseEnrollment)
		return &Sut{Sut: sut, CreateCourseEnrollment: createCourseEnrollment, Input: input}
	}

	s.Run("should create a courseEnrollment", func() {
		sut := makeSut()

		sut.CreateCourseEnrollment.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.Output{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.CreateCourseEnrollment.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid StudentID: it's required", func() {
		sut := makeSut()

		sut.Input.StudentID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid CourseSubjectID: it's required", func() {
		sut := makeSut()

		sut.Input.CourseSubjectID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when CreateCourseEnrollment.Handle returns error", func() {
		sut := makeSut()

		sut.CreateCourseEnrollment.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
