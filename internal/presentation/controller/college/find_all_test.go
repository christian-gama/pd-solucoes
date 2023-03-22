package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/college"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/dto"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/college"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindAllCollegeSuite struct {
	suite.Suite
}

func TestFindAllCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindAllCollegeSuite))
}

func (s *FindAllCollegeSuite) TestHandle() {
	type Sut struct {
		Sut             controller.FindAllColleges
		Input           *dto.FindAllCollegesInput
		FindAllColleges *mocks.FindAllColleges
	}

	makeSut := func() *Sut {
		input := fake.FindAllCollegesInput()
		findAllCollege := new(mocks.FindAllColleges)
		sut := controller.NewFindAllColleges(findAllCollege)
		return &Sut{Sut: sut, FindAllColleges: findAllCollege, Input: input}
	}

	s.Run("should find all colleges", func() {
		sut := makeSut()

		sut.FindAllColleges.On("Handle", mock.Anything, mock.Anything).
			Return(&dto.FindAllCollegesOutput{}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindAllColleges.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("panics when FindAllCollege.Handle returns error", func() {
		sut := makeSut()

		sut.FindAllColleges.On("Handle", mock.Anything, sut.Input).Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
