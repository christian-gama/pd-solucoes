package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneCollege interface {
	// Handle find one college.
	Handle(ctx context.Context, params *dto.FindOneCollegeInput) (*dto.FindOneCollegeOutput, error)
}

type findOneCollegeImpl struct {
	repo.College
}

// NewFindOneCollege findOnes a new CollegeService.
func NewFindOneCollege(collegeRepo repo.College) FindOneCollege {
	return &findOneCollegeImpl{College: collegeRepo}
}

// Handle findOnes a new college.
func (s *findOneCollegeImpl) Handle(
	ctx context.Context,
	input *dto.FindOneCollegeInput,
) (*dto.FindOneCollegeOutput, error) {
	findOneCollegeParams := repo.FindOneCollegeParams{
		ID:       input.ID,
		Filterer: input.Filter,
	}
	college, err := s.College.FindOne(ctx, findOneCollegeParams)
	if err != nil {
		return nil, err
	}

	output := &dto.FindOneCollegeOutput{
		ID:   college.ID,
		Name: college.Name,
		Cnpj: college.Cnpj,
	}

	return output, nil
}
