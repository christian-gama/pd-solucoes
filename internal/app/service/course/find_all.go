package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllCourses interface {
	// Handle finds all courses.
	Handle(
		ctx context.Context,
		input *FindAllCoursesInput,
	) (*querying.PaginationOutput[*model.Course], error)
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
) (*querying.PaginationOutput[*model.Course], error) {
	findAllCourseParams := repo.FindAllCourseParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.Course.FindAll(
		ctx,
		findAllCourseParams,
		"enrollments",
		"subjects.students",
	)
	if err != nil {
		return nil, err
	}

	return paginationOutput, nil
}
