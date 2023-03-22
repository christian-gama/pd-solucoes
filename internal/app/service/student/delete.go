package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteStudent interface {
	// Handle deletes a student.
	Handle(ctx context.Context, input *DeleteStudentInput) error
}

type deleteStudentImpl struct {
	repo.Student
}

// NewDeleteStudent returns a DeleteStudent.
func NewDeleteStudent(studentRepo repo.Student) DeleteStudent {
	return &deleteStudentImpl{Student: studentRepo}
}

// Handle deletes a student.
func (s *deleteStudentImpl) Handle(
	ctx context.Context,
	input *DeleteStudentInput,
) error {
	if _, err := s.Student.FindOne(ctx, repo.FindOneStudentParams{ID: input.ID}); err != nil {
		return err
	}

	deleteStudentParams := repo.DeleteStudentParams{
		ID: input.ID,
	}
	err := s.Student.Delete(ctx, deleteStudentParams)
	if err != nil {
		return err
	}

	return nil
}
