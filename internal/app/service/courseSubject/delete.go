package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteCourseSubject interface {
	// Handle deletes a courseSubject.
	Handle(ctx context.Context, input *DeleteCourseSubjectInput) error
}

type deleteCourseSubjectImpl struct {
	repo.CourseSubject
}

// NewDeleteCourseSubject returns a DeleteCourseSubject.
func NewDeleteCourseSubject(
	courseSubjectRepo repo.CourseSubject,
) DeleteCourseSubject {
	return &deleteCourseSubjectImpl{CourseSubject: courseSubjectRepo}
}

// Handle deletes a courseSubject.
func (s *deleteCourseSubjectImpl) Handle(
	ctx context.Context,
	input *DeleteCourseSubjectInput,
) error {
	if _, err := s.
		CourseSubject.
		FindOne(ctx, repo.FindOneCourseSubjectParams{ID: input.ID}); err != nil {
		return err
	}

	deleteCourseSubjectParams := repo.DeleteCourseSubjectParams{
		ID: input.ID,
	}
	err := s.CourseSubject.Delete(ctx, deleteCourseSubjectParams)
	if err != nil {
		return err
	}

	return nil
}
