package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindOneSubject interface {
	// Handle finds one subject.
	Handle(ctx context.Context, input *FindOneSubjectInput) (*Output, error)
}

type findOneSubjectImpl struct {
	repo.Subject
}

// NewFindOneSubject returns a FindOneSubject.
func NewFindOneSubject(subjectRepo repo.Subject) FindOneSubject {
	return &findOneSubjectImpl{Subject: subjectRepo}
}

// Handle findOnes a new subject.
func (s *findOneSubjectImpl) Handle(
	ctx context.Context,
	input *FindOneSubjectInput,
) (*Output, error) {
	findOneSubjectParams := repo.FindOneSubjectParams{
		ID: input.ID,
	}
	subject, err := s.Subject.FindOne(
		ctx,
		findOneSubjectParams,
		"teacher",
		"courses",
		"courses.students",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, subject), nil
}
