package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllColleges interface {
	// Handle finds all colleges.
	Handle(
		ctx context.Context,
		input *FindAllCollegesInput,
	) (*querying.PaginationOutput[*model.College], error)
}

type findAllCollegesImpl struct {
	repo.College
}

// NewFindAllColleges returns a FindAllColleges.
func NewFindAllColleges(collegeRepo repo.College) FindAllColleges {
	return &findAllCollegesImpl{College: collegeRepo}
}

// Handle findAlls a new college.
func (s *findAllCollegesImpl) Handle(
	ctx context.Context,
	input *FindAllCollegesInput,
) (*querying.PaginationOutput[*model.College], error) {
	findAllCollegeParams := repo.FindAllCollegeParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.College.FindAll(ctx, findAllCollegeParams, "courses")
	if err != nil {
		return nil, err
	}

	return paginationOutput, nil
}
