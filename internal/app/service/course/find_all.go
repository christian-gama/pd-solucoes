package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllCourses interface {
	// Handle finds all courses.
	Handle(ctx context.Context, input *FindAllCoursesInput) (*FindAllCoursesOutput, error)
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
) (*FindAllCoursesOutput, error) {
	findAllCourseParams := repo.FindAllCourseParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	course, err := s.Course.FindAll(ctx, findAllCourseParams)
	if err != nil {
		return nil, err
	}

	result := make([]*FindOneCourseOutput, 0, len(course.Results))
	for _, c := range course.Results {
		result = append(result, &FindOneCourseOutput{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	output := &FindAllCoursesOutput{
		Total:   course.Total,
		Results: result,
	}

	return output, nil
}
