package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type UpdateCourseSubject interface {
	// Handle updates a courseSubject.
	Handle(ctx context.Context, input *UpdateInput) (*UpdateOutput, error)
}

type updateCourseSubjectImpl struct {
	repo.CourseSubject
}

// NewUpdateCourseSubject returns a UpdateCourseSubject.
func NewUpdateCourseSubject(courseSubjectRepo repo.CourseSubject) UpdateCourseSubject {
	return &updateCourseSubjectImpl{
		CourseSubject: courseSubjectRepo,
	}
}

// Handle updates a courseSubject.
func (s *updateCourseSubjectImpl) Handle(
	ctx context.Context,
	input *UpdateInput,
) (*UpdateOutput, error) {
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

	return copy.MustCopy(&UpdateOutput{}, courseSubject), nil
}
