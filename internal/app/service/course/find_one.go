package service

import (
	"context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneCourse interface {
	// Handle finds one course.
	Handle(ctx context.Context, input *FindOneCourseInput) (*FindOneCourseOutput, error)
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
) (*FindOneCourseOutput, error) {
	findOneCourseParams := repo.FindOneCourseParams{
		ID: input.ID,
	}
	course, err := s.Course.FindOne(ctx, findOneCourseParams, "college")
	if err != nil {
		return nil, err
	}

	output := &FindOneCourseOutput{
		ID:      course.ID,
		Name:    course.Name,
		College: (*service.FindOneCollegeOutput)(course.College),
	}

	return output, nil
}
