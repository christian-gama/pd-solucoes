package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneCourse interface {
	// Handle finds one course.
	Handle(ctx context.Context, input *FindOneCourseInput) (*model.Course, error)
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
) (*model.Course, error) {
	findOneCourseParams := repo.FindOneCourseParams{
		ID: input.ID,
	}
	course, err := s.Course.FindOne(
		ctx,
		findOneCourseParams,
		"enrollments",
		"subjects.students",
		"college",
	)
	if err != nil {
		return nil, err
	}

	return course, nil
}
