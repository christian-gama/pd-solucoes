package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindAllColleges interface {
	// Handle finds all colleges.
	Handle(
		ctx context.Context,
		input *FindAllCollegesInput,
	) (*querying.PaginationOutput[*Output], error)
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
) (*querying.PaginationOutput[*Output], error) {
	findAllCollegeParams := repo.FindAllCollegeParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.College.FindAll(ctx, findAllCollegeParams, "courses")
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&querying.PaginationOutput[*Output]{}, paginationOutput), nil
}
