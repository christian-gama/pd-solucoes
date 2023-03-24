package service

import (
	"context"
	"time"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type UpdateCourseEnrollment interface {
	// Handle updates a courseEnrollment.
	Handle(ctx context.Context, input *UpdateCourseEnrollmentInput) (*Output, error)
}

type updateCourseEnrollmentImpl struct {
	repo.CourseEnrollment
	FindOneCourseEnrollment
}

// NewUpdateCourseEnrollment returns a UpdateCourseEnrollment.
func NewUpdateCourseEnrollment(
	courseEnrollmentRepo repo.CourseEnrollment,
	findOneCourseEnrollmentService FindOneCourseEnrollment,
) UpdateCourseEnrollment {
	return &updateCourseEnrollmentImpl{
		CourseEnrollment:        courseEnrollmentRepo,
		FindOneCourseEnrollment: findOneCourseEnrollmentService,
	}
}

// Handle updates a courseEnrollment.
func (s *updateCourseEnrollmentImpl) Handle(
	ctx context.Context,
	input *UpdateCourseEnrollmentInput,
) (*Output, error) {
	courseEnrollment, err := model.NewCourseEnrollment(
		input.ID,
		input.StudentID,
		time.Now(),
		input.CourseSubjectID,
	)
	if err != nil {
		return nil, err
	}

	updateCourseEnrollmentParams := repo.UpdateCourseEnrollmentParams{
		CourseEnrollment: courseEnrollment,
	}
	courseEnrollment, err = s.CourseEnrollment.Update(ctx, updateCourseEnrollmentParams)
	if err != nil {
		return nil, err
	}

	findOneCourseEnrollmentParams := &FindOneCourseEnrollmentInput{
		ID: courseEnrollment.ID,
	}
	output, err := s.FindOneCourseEnrollment.Handle(ctx, findOneCourseEnrollmentParams)
	if err != nil {
		return nil, err
	}

	return output, err
}
