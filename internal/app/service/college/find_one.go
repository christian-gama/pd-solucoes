package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneCollege interface {
	// Handle finds one college.
	Handle(
		ctx context.Context,
		input *FindOneCollegeInput,
	) (*model.College, error)
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
) (*model.College, error) {
	findOneCollegeParams := repo.FindOneCollegeParams{
		ID: input.ID,
	}
	course, err := s.College.FindOne(ctx, findOneCollegeParams, "courses")
	if err != nil {
		return nil, err
	}

	return course, nil
}
