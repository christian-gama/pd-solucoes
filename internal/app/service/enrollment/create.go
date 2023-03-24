package service

import (
	"context"
	"time"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type CreateCourseEnrollment interface {
	// Handle creates a new courseEnrollment.
	Handle(ctx context.Context, input *CreateInput) (*CreateOutput, error)
}

type createCourseEnrollmentImpl struct {
	repo.CourseEnrollment
}

// NewCreateCourseEnrollment returns a CreateCourseEnrollment.
func NewCreateCourseEnrollment(courseEnrollmentRepo repo.CourseEnrollment,
) CreateCourseEnrollment {
	return &createCourseEnrollmentImpl{CourseEnrollment: courseEnrollmentRepo}
}

// Handle creates a new courseEnrollment.
func (s *createCourseEnrollmentImpl) Handle(
	ctx context.Context,
	input *CreateInput,
) (*CreateOutput, error) {
	courseEnrollment, err := model.NewCourseEnrollment(
		0,
		input.StudentID,
		time.Now(),
		input.CourseSubjectID,
	)
	if err != nil {
		return nil, err
	}

	createCourseEnrollmentParams := repo.CreateCourseEnrollmentParams{
		CourseEnrollment: courseEnrollment,
	}
	courseEnrollment, err = s.CourseEnrollment.Create(ctx, createCourseEnrollmentParams)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&CreateOutput{}, courseEnrollment), nil
}
