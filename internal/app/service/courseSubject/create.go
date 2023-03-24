package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type CreateCourseSubject interface {
	// Handle creates a new courseSubject.
	Handle(ctx context.Context, input *CreateCourseSubjectInput) (*Output, error)
}

type createCourseSubjectImpl struct {
	repo.CourseSubject
	FindOneCourseSubject
}

// NewCreateCourseSubject returns a CreateCourseSubject.
func NewCreateCourseSubject(
	courseSubjectRepo repo.CourseSubject,
	findOneCourseSubjectService FindOneCourseSubject,
) CreateCourseSubject {
	return &createCourseSubjectImpl{
		CourseSubject:        courseSubjectRepo,
		FindOneCourseSubject: findOneCourseSubjectService,
	}
}

// Handle creates a new courseSubject.
func (s *createCourseSubjectImpl) Handle(
	ctx context.Context,
	input *CreateCourseSubjectInput,
) (*Output, error) {
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

	findOneCourseSubjectParams := &FindOneCourseSubjectInput{
		ID: courseSubject.ID,
	}
	output, err := s.FindOneCourseSubject.Handle(ctx, findOneCourseSubjectParams)
	if err != nil {
		return nil, err
	}

	return output, err
}
