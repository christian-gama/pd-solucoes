package service

import (
	"context"
	"time"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type CreateCourseEnrollment interface {
	// Handle creates a new courseEnrollment.
	Handle(ctx context.Context, input *CreateCourseEnrollmentInput) (*Output, error)
}

type createCourseEnrollmentImpl struct {
	repo.CourseEnrollment
	FindOneCourseEnrollment
}

// NewCreateCourseEnrollment returns a CreateCourseEnrollment.
func NewCreateCourseEnrollment(
	courseEnrollmentRepo repo.CourseEnrollment,
	findOneCourseEnrollmentService FindOneCourseEnrollment,
) CreateCourseEnrollment {
	return &createCourseEnrollmentImpl{
		CourseEnrollment:        courseEnrollmentRepo,
		FindOneCourseEnrollment: findOneCourseEnrollmentService,
	}
}

// Handle creates a new courseEnrollment.
func (s *createCourseEnrollmentImpl) Handle(
	ctx context.Context,
	input *CreateCourseEnrollmentInput,
) (*Output, error) {
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

	findOneCourseEnrollmentParams := &FindOneCourseEnrollmentInput{
		ID: courseEnrollment.ID,
	}
	output, err := s.FindOneCourseEnrollment.Handle(ctx, findOneCourseEnrollmentParams)
	if err != nil {
		return nil, err
	}

	return output, err
}
