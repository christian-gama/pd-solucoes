package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type UpdateSubject interface {
	// Handle updates a subject.
	Handle(ctx context.Context, input *UpdateInput) (*Output, error)
}

type updateSubjectImpl struct {
	repo.Subject
}

// NewUpdateSubject returns a UpdateSubject.
func NewUpdateSubject(subjectRepo repo.Subject) UpdateSubject {
	return &updateSubjectImpl{Subject: subjectRepo}
}

// Handle updates a subject.
func (s *updateSubjectImpl) Handle(
	ctx context.Context,
	input *UpdateInput,
) (*Output, error) {
	subject, err := model.NewSubject(input.ID, input.Name, input.TeacherID)
	if err != nil {
		return nil, err
	}

	updateSubjectParams := repo.UpdateSubjectParams{
		Subject: subject,
	}
	subject, err = s.Subject.Update(ctx, updateSubjectParams)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, subject), nil
}
