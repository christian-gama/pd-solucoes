package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteCourseEnrollment interface {
	// Handle deletes a courseEnrollment.
	Handle(ctx context.Context, input *DeleteInput) error
}

type deleteCourseEnrollmentImpl struct {
	repo.CourseEnrollment
}

// NewDeleteCourseEnrollment returns a DeleteCourseEnrollment.
func NewDeleteCourseEnrollment(courseEnrollmentRepo repo.CourseEnrollment) DeleteCourseEnrollment {
	return &deleteCourseEnrollmentImpl{CourseEnrollment: courseEnrollmentRepo}
}

// Handle deletes a courseEnrollment.
func (s *deleteCourseEnrollmentImpl) Handle(
	ctx context.Context,
	input *DeleteInput,
) error {
	if _, err := s.
		CourseEnrollment.
		FindOne(ctx, repo.FindOneCourseEnrollmentParams{ID: input.ID}); err != nil {
		return err
	}

	deleteCourseEnrollmentParams := repo.DeleteCourseEnrollmentParams{
		ID: input.ID,
	}
	err := s.CourseEnrollment.Delete(ctx, deleteCourseEnrollmentParams)
	if err != nil {
		return err
	}

	return nil
}
