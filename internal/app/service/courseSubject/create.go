package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type CreateCourseSubject interface {
	// Handle creates a new courseSubject.
	Handle(ctx context.Context, input *CreateInput) (*CreateOutput, error)
}

type createCourseSubjectImpl struct {
	repo.CourseSubject
}

// NewCreateCourseSubject returns a CreateCourseSubject.
func NewCreateCourseSubject(courseSubjectRepo repo.CourseSubject) CreateCourseSubject {
	return &createCourseSubjectImpl{
		CourseSubject: courseSubjectRepo,
	}
}

// Handle creates a new courseSubject.
func (s *createCourseSubjectImpl) Handle(
	ctx context.Context,
	input *CreateInput,
) (*CreateOutput, error) {
	courseSubject, err := model.NewCourseSubject(
		0,
		input.CourseID,
		input.SubjectID,
	)
	if err != nil {
		return nil, err
	}

	createCourseSubjectParams := repo.CreateCourseSubjectParams{
		CourseSubject: courseSubject,
	}
	courseSubject, err = s.CourseSubject.Create(ctx, createCourseSubjectParams)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&CreateOutput{}, courseSubject), nil
}
