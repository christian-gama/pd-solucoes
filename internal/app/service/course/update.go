package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type UpdateCourse interface {
	// Handle updates a course.
	Handle(ctx context.Context, input *UpdateCourseInput) (*UpdateCourseOutput, error)
}

type updateCourseImpl struct {
	repo.Course
}

// NewUpdateCourse returns a UpdateCourse.
func NewUpdateCourse(courseRepo repo.Course) UpdateCourse {
	return &updateCourseImpl{Course: courseRepo}
}

// Handle updates a course.
func (s *updateCourseImpl) Handle(
	ctx context.Context,
	input *UpdateCourseInput,
) (*UpdateCourseOutput, error) {
	course, err := model.NewCourse(input.ID, input.Name, input.CollegeID, nil)
	if err != nil {
		return nil, err
	}

	updateCourseParams := repo.UpdateCourseParams{
		Course: course,
	}
	course, err = s.Course.Update(ctx, updateCourseParams)
	if err != nil {
		return nil, err
	}

	output := &UpdateCourseOutput{
		ID:        course.ID,
		Name:      course.Name,
		CollegeID: course.CollegeID,
	}

	return output, nil
}
