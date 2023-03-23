package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindAllSubjects interface {
	// Handle finds all subjects.
	Handle(
		ctx context.Context,
		input *FindAllSubjectsInput,
	) (*querying.PaginationOutput[*Output], error)
}

type findAllSubjectsImpl struct {
	repo.Subject
}

// NewFindAllSubjects returns a FindAllSubjects.
func NewFindAllSubjects(subjectRepo repo.Subject) FindAllSubjects {
	return &findAllSubjectsImpl{Subject: subjectRepo}
}

// Handle findAlls a new subject.
func (s *findAllSubjectsImpl) Handle(
	ctx context.Context,
	input *FindAllSubjectsInput,
) (*querying.PaginationOutput[*Output], error) {
	findAllSubjectParams := repo.FindAllSubjectParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.Subject.FindAll(
		ctx,
		findAllSubjectParams,
		"teacher",
		"courses",
		"courses.students",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&querying.PaginationOutput[*Output]{}, paginationOutput), nil
}
