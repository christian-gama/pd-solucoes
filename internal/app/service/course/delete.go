package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteCourse interface {
	// Handle deletes a course.
	Handle(ctx context.Context, input *DeleteInput) error
}

type deleteCourseImpl struct {
	repo.Course
}

// NewDeleteCourse returns a DeleteCourse.
func NewDeleteCourse(courseRepo repo.Course) DeleteCourse {
	return &deleteCourseImpl{Course: courseRepo}
}

// Handle deletes a course.
func (s *deleteCourseImpl) Handle(
	ctx context.Context,
	input *DeleteInput,
) error {
	if _, err := s.Course.FindOne(ctx, repo.FindOneCourseParams{ID: input.ID}); err != nil {
		return err
	}

	deleteCourseParams := repo.DeleteCourseParams{
		ID: input.ID,
	}
	err := s.Course.Delete(ctx, deleteCourseParams)
	if err != nil {
		return err
	}

	return nil
}
