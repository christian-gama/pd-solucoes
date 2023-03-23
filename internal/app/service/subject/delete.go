package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteSubject interface {
	// Handle deletes a subject.
	Handle(ctx context.Context, input *DeleteSubjectInput) error
}

type deleteSubjectImpl struct {
	repo.Subject
}

// NewDeleteSubject returns a DeleteSubject.
func NewDeleteSubject(subjectRepo repo.Subject) DeleteSubject {
	return &deleteSubjectImpl{Subject: subjectRepo}
}

// Handle deletes a subject.
func (s *deleteSubjectImpl) Handle(
	ctx context.Context,
	input *DeleteSubjectInput,
) error {
	if _, err := s.Subject.FindOne(ctx, repo.FindOneSubjectParams{ID: input.ID}); err != nil {
		return err
	}

	deleteSubjectParams := repo.DeleteSubjectParams{
		ID: input.ID,
	}
	err := s.Subject.Delete(ctx, deleteSubjectParams)
	if err != nil {
		return err
	}

	return nil
}
