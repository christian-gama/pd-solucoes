package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type DeleteTeacher interface {
	// Handle deletes a teacher.
	Handle(ctx context.Context, input *DeleteTeacherInput) error
}

type deleteTeacherImpl struct {
	repo.Teacher
}

// NewDeleteTeacher returns a DeleteTeacher.
func NewDeleteTeacher(teacherRepo repo.Teacher) DeleteTeacher {
	return &deleteTeacherImpl{Teacher: teacherRepo}
}

// Handle deletes a teacher.
func (s *deleteTeacherImpl) Handle(
	ctx context.Context,
	input *DeleteTeacherInput,
) error {
	if _, err := s.Teacher.FindOne(ctx, repo.FindOneTeacherParams{ID: input.ID}); err != nil {
		return err
	}

	deleteTeacherParams := repo.DeleteTeacherParams{
		ID: input.ID,
	}
	err := s.Teacher.Delete(ctx, deleteTeacherParams)
	if err != nil {
		return err
	}

	return nil
}
