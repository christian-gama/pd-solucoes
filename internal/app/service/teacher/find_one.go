package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindOneTeacher interface {
	// Handle finds one teacher.
	Handle(ctx context.Context, input *FindOneTeacherInput) (*FindOneTeacherOutput, error)
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
) (*FindOneTeacherOutput, error) {
	findOneTeacherParams := repo.FindOneTeacherParams{
		ID: input.ID,
	}
	teacher, err := s.Teacher.FindOne(ctx, findOneTeacherParams)
	if err != nil {
		return nil, err
	}

	output := &FindOneTeacherOutput{
		ID:     teacher.ID,
		Name:   teacher.Name,
		Degree: teacher.Degree,
	}

	return output, nil
}
