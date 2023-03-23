package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindOneStudent interface {
	// Handle finds one student.
	Handle(ctx context.Context, input *FindOneStudentInput) (*Output, error)
}

type findOneStudentImpl struct {
	repo.Student
}

// NewFindOneStudent returns a FindOneStudent.
func NewFindOneStudent(studentRepo repo.Student) FindOneStudent {
	return &findOneStudentImpl{Student: studentRepo}
}

// Handle findOnes a new student.
func (s *findOneStudentImpl) Handle(
	ctx context.Context,
	input *FindOneStudentInput,
) (*Output, error) {
	findOneStudentParams := repo.FindOneStudentParams{
		ID: input.ID,
	}
	student, err := s.Student.FindOne(
		ctx,
		findOneStudentParams,
		"courseSubjects",
		"courseSubjects.subject",
		"courseSubjects.course",
	)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, student), nil
}
