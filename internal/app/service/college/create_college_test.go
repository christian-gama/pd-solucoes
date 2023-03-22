package service_test

import (
	"context"
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/dto"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CreateCollegeSuite struct {
	suite.Suite
}

func TestCreateCollegeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CreateCollegeSuite))
}

func (s *CreateCollegeSuite) TestHandle() {
	type Sut struct {
		Sut         service.CreateCollege
		CollegeRepo *mocks.College
		Input       *dto.CreateCollegeInput
		College     *model.College
	}

	makeSut := func() *Sut {
		collegeRepo := mocks.NewCollege(s.T())
		sut := service.NewCreateCollege(collegeRepo)

		return &Sut{
			Sut:         sut,
			CollegeRepo: collegeRepo,
			Input:       fake.CreateCollegeInput(),
			College:     fakeModel.College(),
		}
	}

	s.Run("should add a new college", func() {
		sut := makeSut()

		sut.CollegeRepo.On("Create", mock.Anything, mock.Anything).Return(sut.College, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.College.ID, result.ID)
		s.Equal(sut.College.Name, result.Name)
		s.Equal(sut.College.Cnpj, result.Cnpj)
	})

	s.Run("collegeRepo.Create returns an error", func() {
		sut := makeSut()

		sut.CollegeRepo.On("Create", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})

	s.Run("should return an error if domain validation fails", func() {
		sut := makeSut()

		sut.Input.Name = ""

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.Error(err)
		s.Nil(result)
	})
}
