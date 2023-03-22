package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/app/dto"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type UpdateCollege interface {
	// Handle updates a college.
	Handle(ctx context.Context, params *dto.UpdateCollegeInput) (*dto.UpdateCollegeOutput, error)
}

type updateCollegeImpl struct {
	repo.College
}

// NewUpdateCollege returns a UpdateCollege.
func NewUpdateCollege(collegeRepo repo.College) UpdateCollege {
	return &updateCollegeImpl{College: collegeRepo}
}

// Handle updates a college.
func (s *updateCollegeImpl) Handle(
	ctx context.Context,
	input *dto.UpdateCollegeInput,
) (*dto.UpdateCollegeOutput, error) {
	college, err := model.NewCollege(input.ID, input.Name, input.Cnpj)
	if err != nil {
		return nil, err
	}

	updateCollegeParams := repo.UpdateCollegeParams{
		College: college,
	}
	college, err = s.College.Update(ctx, updateCollegeParams)
	if err != nil {
		return nil, err
	}

	output := &dto.UpdateCollegeOutput{
		ID:   college.ID,
		Name: college.Name,
		Cnpj: college.Cnpj,
	}

	return output, nil
}
