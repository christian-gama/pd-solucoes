package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/college"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
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
		Sut         service.FindAllColleges
		CollegeRepo *mocks.College
		Input       *service.FindAllCollegesInput
		Pagination  *querying.PaginationOutput[*model.College]
	}

	makeSut := func() *Sut {
		collegeRepo := mocks.NewCollege(s.T())
		sut := service.NewFindAllColleges(collegeRepo)

		return &Sut{
			Sut:         sut,
			CollegeRepo: collegeRepo,
			Input:       fake.FindAllCollegesInput(),
			Pagination: &querying.PaginationOutput[*model.College]{
				Total:   100,
				Results: []*model.College{fakeModel.College()},
			},
		}
	}

	s.Run("should find one college", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindAll", mock.Anything, mock.Anything).Return(sut.Pagination, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.Pagination.Total, result.Total)
		s.Equal(len(sut.Pagination.Results), len(result.Results))
	})

	s.Run("collegeRepo.FindAll returns an error", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
