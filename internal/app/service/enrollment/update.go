package service

import (
	"context"
	"time"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type UpdateCourseEnrollment interface {
	// Handle updates a courseEnrollment.
	Handle(ctx context.Context, input *UpdateInput) (*UpdateOutput, error)
}

type updateCourseEnrollmentImpl struct {
	repo.CourseEnrollment
}

// NewUpdateCourseEnrollment returns a UpdateCourseEnrollment.
func NewUpdateCourseEnrollment(courseEnrollmentRepo repo.CourseEnrollment) UpdateCourseEnrollment {
	return &updateCourseEnrollmentImpl{CourseEnrollment: courseEnrollmentRepo}
}

// Handle updates a courseEnrollment.
func (s *updateCourseEnrollmentImpl) Handle(
	ctx context.Context,
	input *UpdateInput,
) (*UpdateOutput, error) {
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

	return copy.MustCopy(&UpdateOutput{}, courseEnrollment), nil
}
