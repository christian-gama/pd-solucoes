package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteCollege interface {
	// Handle deletes a college.
	Handle(ctx context.Context, input *DeleteInput) error
}

type deleteCollegeImpl struct {
	repo.College
}

// NewDeleteCollege returns a DeleteCollege.
func NewDeleteCollege(collegeRepo repo.College) DeleteCollege {
	return &deleteCollegeImpl{College: collegeRepo}
}

// Handle deletes a college.
func (s *deleteCollegeImpl) Handle(
	ctx context.Context,
	input *DeleteInput,
) error {
	if _, err := s.College.FindOne(ctx, repo.FindOneCollegeParams{ID: input.ID}); err != nil {
		return err
	}

	deleteCollegeParams := repo.DeleteCollegeParams{
		ID: input.ID,
	}
	err := s.College.Delete(ctx, deleteCollegeParams)
	if err != nil {
		return err
	}

	return nil
}
