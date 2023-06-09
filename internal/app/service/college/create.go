package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type CreateCollege interface {
	// Handle creates a new college.
	Handle(ctx context.Context, input *CreateInput) (*CreateOutput, error)
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
	input *CreateInput,
) (*CreateOutput, error) {
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

	return copy.MustCopy(&CreateOutput{}, college), nil
}
