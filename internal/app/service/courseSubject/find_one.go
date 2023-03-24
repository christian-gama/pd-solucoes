package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindOneCourseSubject interface {
	// Handle finds one courseSubject.
	Handle(ctx context.Context, input *FindOneInput) (*Output, error)
}

type findOneCourseSubjectImpl struct {
	repo.CourseSubject
}

// NewFindOneCourseSubject returns a FindOneCourseSubject.
func NewFindOneCourseSubject(
	courseSubjectRepo repo.CourseSubject,
) FindOneCourseSubject {
	return &findOneCourseSubjectImpl{CourseSubject: courseSubjectRepo}
}

// Handle findOnes a new courseSubject.
func (s *findOneCourseSubjectImpl) Handle(
	ctx context.Context,
	input *FindOneInput,
) (*Output, error) {
	findOneCourseSubjectParams := repo.FindOneCourseSubjectParams{
		ID: input.ID,
	}
	courseSubject, err := s.CourseSubject.FindOne(
		ctx,
		findOneCourseSubjectParams,
		"subject",
		"course",
		"students",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, courseSubject), nil
}
