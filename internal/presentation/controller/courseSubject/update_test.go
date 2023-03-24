package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/courseSubject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateCourseSubjectSuite struct {
	suite.Suite
}

func TestUpdateCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateCourseSubjectSuite))
}

func (s *UpdateCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut                 controller.UpdateCourseSubject
		Input               *service.UpdateInput
		UpdateCourseSubject *mocks.UpdateCourseSubject
	}

	makeSut := func() *Sut {
		input := fake.UpdateCourseSubjectInput()
		updateCourseSubject := new(mocks.UpdateCourseSubject)
		sut := controller.NewUpdateCourseSubject(updateCourseSubject)
		return &Sut{Sut: sut, UpdateCourseSubject: updateCourseSubject, Input: input}
	}

	s.Run("should update a courseSubject", func() {
		sut := makeSut()

		sut.UpdateCourseSubject.
			On("Handle", mock.Anything, sut.Input).
			Return(&service.UpdateOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.UpdateCourseSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid CourseID: it's required", func() {
		sut := makeSut()

		sut.Input.CourseID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid SubjectID: it's required", func() {
		sut := makeSut()

		sut.Input.SubjectID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when UpdateCourseSubject.Handle returns error", func() {
		sut := makeSut()

		sut.UpdateCourseSubject.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
