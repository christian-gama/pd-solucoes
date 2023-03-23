package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/college"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindOneCollegeSuite struct {
	suite.Suite
}

func TestFindOneCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindOneCollegeSuite))
}

func (s *FindOneCollegeSuite) TestHandle() {
	type Sut struct {
		Sut         service.FindOneCollege
		CollegeRepo *mocks.College
		Input       *service.FindOneCollegeInput
		College     *model.College
	}

	makeSut := func() *Sut {
		collegeRepo := mocks.NewCollege(s.T())
		sut := service.NewFindOneCollege(collegeRepo)

		return &Sut{
			Sut:         sut,
			CollegeRepo: collegeRepo,
			Input:       fake.FindOneCollegeInput(),
			College:     fakeModel.College(),
		}
	}

	s.Run("should find one college", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindOne", mock.Anything, mock.Anything, "courses").Return(sut.College, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.College.ID, result.ID)
		s.Equal(sut.College.Name, result.Name)
		s.Equal(sut.College.Cnpj, result.Cnpj)
	})

	s.Run("collegeRepo.FindOne returns an error", func() {
		sut := makeSut()

		sut.CollegeRepo.On("FindOne", mock.Anything, mock.Anything, "courses").
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})
}
