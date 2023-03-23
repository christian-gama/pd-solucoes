package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindAllCourses interface {
	// Handle finds all courses.
	Handle(
		ctx context.Context,
		input *FindAllCoursesInput,
	) (*querying.PaginationOutput[*Output], error)
}

type findAllCoursesImpl struct {
	repo.Course
}

// NewFindAllCourses returns a FindAllCourses.
func NewFindAllCourses(courseRepo repo.Course) FindAllCourses {
	return &findAllCoursesImpl{Course: courseRepo}
}

// Handle findAlls a new course.
func (s *findAllCoursesImpl) Handle(
	ctx context.Context,
	input *FindAllCoursesInput,
) (*querying.PaginationOutput[*Output], error) {
	findAllCourseParams := repo.FindAllCourseParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.Course.FindAll(
		ctx,
		findAllCourseParams,
		"subjects.students",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&querying.PaginationOutput[*Output]{}, paginationOutput), nil
}
