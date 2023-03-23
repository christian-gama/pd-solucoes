package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindOneTeacher interface {
	// Handle finds one teacher.
	Handle(ctx context.Context, input *FindOneTeacherInput) (*Output, error)
}

type findOneTeacherImpl struct {
	repo.Teacher
}

// NewFindOneTeacher returns a FindOneTeacher.
func NewFindOneTeacher(teacherRepo repo.Teacher) FindOneTeacher {
	return &findOneTeacherImpl{Teacher: teacherRepo}
}

// Handle findOnes a new teacher.
func (s *findOneTeacherImpl) Handle(
	ctx context.Context,
	input *FindOneTeacherInput,
) (*Output, error) {
	findOneTeacherParams := repo.FindOneTeacherParams{
		ID: input.ID,
	}
	teacher, err := s.Teacher.FindOne(ctx, findOneTeacherParams, "subjects")
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, teacher), nil
}
