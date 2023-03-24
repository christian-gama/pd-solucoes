package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindAllCourseEnrollments interface {
	// Handle finds all courseEnrollments.
	Handle(
		ctx context.Context,
		input *FindAllInput,
	) (*querying.PaginationOutput[*Output], error)
}

type findAllCourseEnrollmentsImpl struct {
	repo.CourseEnrollment
}

// NewFindAllCourseEnrollments returns a FindAllCourseEnrollments.
func NewFindAllCourseEnrollments(
	courseEnrollmentRepo repo.CourseEnrollment,
) FindAllCourseEnrollments {
	return &findAllCourseEnrollmentsImpl{CourseEnrollment: courseEnrollmentRepo}
}

// Handle findAlls a new courseEnrollment.
func (s *findAllCourseEnrollmentsImpl) Handle(
	ctx context.Context,
	input *FindAllInput,
) (*querying.PaginationOutput[*Output], error) {
	findAllCourseEnrollmentParams := repo.FindAllCourseEnrollmentParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.CourseEnrollment.FindAll(
		ctx,
		findAllCourseEnrollmentParams,
		"courseSubject",
		"courseSubject.subject",
		"courseSubject.course",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&querying.PaginationOutput[*Output]{}, paginationOutput), nil
}
