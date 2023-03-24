package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindAllCourseSubjects interface {
	// Handle finds all courseSubjects.
	Handle(
		ctx context.Context,
		input *FindAllCourseSubjectsInput,
	) (*querying.PaginationOutput[*Output], error)
}

type findAllCourseSubjectsImpl struct {
	repo.CourseSubject
}

// NewFindAllCourseSubjects returns a FindAllCourseSubjects.
func NewFindAllCourseSubjects(
	courseSubjectRepo repo.CourseSubject,
) FindAllCourseSubjects {
	return &findAllCourseSubjectsImpl{CourseSubject: courseSubjectRepo}
}

// Handle findAlls a new courseSubject.
func (s *findAllCourseSubjectsImpl) Handle(
	ctx context.Context,
	input *FindAllCourseSubjectsInput,
) (*querying.PaginationOutput[*Output], error) {
	findAllCourseSubjectParams := repo.FindAllCourseSubjectParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.CourseSubject.FindAll(
		ctx,
		findAllCourseSubjectParams,
		"subject",
		"course",
		"students",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&querying.PaginationOutput[*Output]{}, paginationOutput), nil
}
