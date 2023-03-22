package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneCollege interface {
	// Handle finds one college.
	Handle(ctx context.Context, input *FindOneCollegeInput) (*FindOneCollegeOutput, error)
}

type findOneCollegeImpl struct {
	repo.College
}

// NewFindOneCollege returns a FindOneCollege.
func NewFindOneCollege(collegeRepo repo.College) FindOneCollege {
	return &findOneCollegeImpl{College: collegeRepo}
}

// Handle findOnes a new college.
func (s *findOneCollegeImpl) Handle(
	ctx context.Context,
	input *FindOneCollegeInput,
) (*FindOneCollegeOutput, error) {
	findOneCollegeParams := repo.FindOneCollegeParams{
		ID: input.ID,
	}
	college, err := s.College.FindOne(ctx, findOneCollegeParams)
	if err != nil {
		return nil, err
	}

	output := &FindOneCollegeOutput{
		ID:   college.ID,
		Name: college.Name,
		Cnpj: college.Cnpj,
	}

	return output, nil
}
