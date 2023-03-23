package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindOneCourse interface {
	// Handle finds one course.
	Handle(ctx context.Context, input *FindOneCourseInput) (*Output, error)
}

type findOneCourseImpl struct {
	repo.Course
}

// NewFindOneCourse returns a FindOneCourse.
func NewFindOneCourse(courseRepo repo.Course) FindOneCourse {
	return &findOneCourseImpl{Course: courseRepo}
}

// Handle findOnes a new course.
func (s *findOneCourseImpl) Handle(
	ctx context.Context,
	input *FindOneCourseInput,
) (*Output, error) {
	findOneCourseParams := repo.FindOneCourseParams{
		ID: input.ID,
	}
	course, err := s.Course.FindOne(
		ctx,
		findOneCourseParams,
		"subjects.students",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, course), nil
}
