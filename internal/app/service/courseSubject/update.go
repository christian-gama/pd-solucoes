package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type UpdateCourseSubject interface {
	// Handle updates a courseSubject.
	Handle(ctx context.Context, input *UpdateCourseSubjectInput) (*Output, error)
}

type updateCourseSubjectImpl struct {
	repo.CourseSubject
	FindOneCourseSubject
}

// NewUpdateCourseSubject returns a UpdateCourseSubject.
func NewUpdateCourseSubject(
	courseSubjectRepo repo.CourseSubject,
	findOneCourseSubjectService FindOneCourseSubject,
) UpdateCourseSubject {
	return &updateCourseSubjectImpl{
		CourseSubject:        courseSubjectRepo,
		FindOneCourseSubject: findOneCourseSubjectService,
	}
}

// Handle updates a courseSubject.
func (s *updateCourseSubjectImpl) Handle(
	ctx context.Context,
	input *UpdateCourseSubjectInput,
) (*Output, error) {
	courseSubject, err := model.NewCourseSubject(
		input.ID,
		input.CourseID,
		input.SubjectID,
	)
	if err != nil {
		return nil, err
	}

	updateCourseSubjectParams := repo.UpdateCourseSubjectParams{
		CourseSubject: courseSubject,
	}
	courseSubject, err = s.CourseSubject.Update(ctx, updateCourseSubjectParams)
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
