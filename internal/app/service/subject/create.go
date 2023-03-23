package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type CreateSubject interface {
	// Handle creates a new subject.
	Handle(ctx context.Context, input *CreateSubjectInput) (*Output, error)
}

type createSubjectImpl struct {
	repo.Subject
}

// NewCreateSubject returns a CreateSubject.
func NewCreateSubject(subjectRepo repo.Subject) CreateSubject {
	return &createSubjectImpl{Subject: subjectRepo}
}

// Handle creates a new subject.
func (s *createSubjectImpl) Handle(
	ctx context.Context,
	input *CreateSubjectInput,
) (*Output, error) {
	subject, err := model.NewSubject(0, input.Name, input.TeacherID)
	if err != nil {
		return nil, err
	}

	createSubjectParams := repo.CreateSubjectParams{
		Subject: subject,
	}
	subject, err = s.Subject.Create(ctx, createSubjectParams)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, subject), nil
}
