package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneStudent interface {
	// Handle finds one student.
	Handle(ctx context.Context, input *FindOneStudentInput) (*FindOneStudentOutput, error)
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
) (*FindOneStudentOutput, error) {
	findOneStudentParams := repo.FindOneStudentParams{
		ID: input.ID,
	}
	student, err := s.Student.FindOne(ctx, findOneStudentParams)
	if err != nil {
		return nil, err
	}

	output := &FindOneStudentOutput{
		ID:   student.ID,
		Name: student.Name,
		Cpf:  student.Cpf,
	}

	return output, nil
}
