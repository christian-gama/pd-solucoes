package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindOneCourseEnrollment interface {
	// Handle finds one courseEnrollment.
	Handle(ctx context.Context, input *FindOneInput) (*Output, error)
}

type findOneCourseEnrollmentImpl struct {
	repo.CourseEnrollment
}

// NewFindOneCourseEnrollment returns a FindOneCourseEnrollment.
func NewFindOneCourseEnrollment(
	courseEnrollmentRepo repo.CourseEnrollment,
) FindOneCourseEnrollment {
	return &findOneCourseEnrollmentImpl{CourseEnrollment: courseEnrollmentRepo}
}

// Handle findOnes a new courseEnrollment.
func (s *findOneCourseEnrollmentImpl) Handle(
	ctx context.Context,
	input *FindOneInput,
) (*Output, error) {
	findOneCourseEnrollmentParams := repo.FindOneCourseEnrollmentParams{
		ID: input.ID,
	}
	courseEnrollment, err := s.CourseEnrollment.FindOne(
		ctx,
		findOneCourseEnrollmentParams,
		"courseSubject",
		"courseSubject.subject",
		"courseSubject.course",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, courseEnrollment), nil
}
