package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type CreateCourse interface {
	// Handle creates a new course.
	Handle(ctx context.Context, input *CreateCourseInput) (*model.Course, error)
}

type createCourseImpl struct {
	repo.Course
}

// NewCreateCourse returns a CreateCourse.
func NewCreateCourse(courseRepo repo.Course) CreateCourse {
	return &createCourseImpl{Course: courseRepo}
}

// Handle creates a new course.
func (s *createCourseImpl) Handle(
	ctx context.Context,
	input *CreateCourseInput,
) (*model.Course, error) {
	course, err := model.NewCourse(0, input.Name, input.CollegeID)
	if err != nil {
		return nil, err
	}

	createCourseParams := repo.CreateCourseParams{
		Course: course,
	}
	course, err = s.Course.Create(ctx, createCourseParams)
	if err != nil {
		return nil, err
	}

	return course, nil
}
