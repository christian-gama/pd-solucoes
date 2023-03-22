package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/college"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCollegeSuite struct {
	suite.Suite
}

func TestDeleteCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCollegeSuite))
}

func (s *DeleteCollegeSuite) TestHandle() {
	type Sut struct {
		Sut         service.DeleteCollege
		CollegeRepo *mocks.College
		Input       *service.DeleteCollegeInput
		College     *model.College
	}

	makeSut := func() *Sut {
		collegeRepo := mocks.NewCollege(s.T())
		sut := service.NewDeleteCollege(collegeRepo)

		return &Sut{
			Sut:         sut,
			CollegeRepo: collegeRepo,
			Input:       fake.DeleteCollegeInput(),
		}
	}

	s.Run("should find one college", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.College, nil)
		sut.CollegeRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
	})

	s.Run("collegeRepo.Delete returns an error", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindOne", mock.Anything, mock.Anything).Return(sut.College, nil)
		sut.CollegeRepo.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("collegeRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
