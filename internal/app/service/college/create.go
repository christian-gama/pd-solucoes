package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type CreateCollege interface {
	// Handle creates a new college.
	Handle(ctx context.Context, input *CreateCollegeInput) (*CreateCollegeOutput, error)
}

type createCollegeImpl struct {
	repo.College
}

// NewCreateCollege returns a CreateCollege.
func NewCreateCollege(collegeRepo repo.College) CreateCollege {
	return &createCollegeImpl{College: collegeRepo}
}

// Handle creates a new college.
func (s *createCollegeImpl) Handle(
	ctx context.Context,
	input *CreateCollegeInput,
) (*CreateCollegeOutput, error) {
	college, err := model.NewCollege(0, input.Name, input.Cnpj)
	if err != nil {
		return nil, err
	}

	createCollegeParams := repo.CreateCollegeParams{
		College: college,
	}
	college, err = s.College.Create(ctx, createCollegeParams)
	if err != nil {
		return nil, err
	}

	output := &CreateCollegeOutput{
		ID:   college.ID,
		Name: college.Name,
		Cnpj: college.Cnpj,
	}

	return output, nil
}
