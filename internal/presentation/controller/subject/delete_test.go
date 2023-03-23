package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/subject"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/subject"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/subject"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteSubjectSuite struct {
	suite.Suite
}

func TestDeleteSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteSubjectSuite))
}

func (s *DeleteSubjectSuite) TestHandle() {
	type Sut struct {
		Sut           controller.DeleteSubject
		Input         *service.DeleteSubjectInput
		DeleteSubject *mocks.DeleteSubject
	}

	makeSut := func() *Sut {
		input := fake.DeleteSubjectInput()
		deleteSubject := new(mocks.DeleteSubject)
		sut := controller.NewDeleteSubject(deleteSubject)
		return &Sut{Sut: sut, DeleteSubject: deleteSubject, Input: input}
	}

	s.Run("should find one subject", func() {
		sut := makeSut()

		sut.DeleteSubject.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteSubject.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteSubject.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteSubject.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
