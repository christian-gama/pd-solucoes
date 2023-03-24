package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type CreateTeacher interface {
	// Handle creates a new teacher.
	Handle(ctx context.Context, input *CreateInput) (*Output, error)
}

type createTeacherImpl struct {
	repo.Teacher
}

// NewCreateTeacher returns a CreateTeacher.
func NewCreateTeacher(teacherRepo repo.Teacher) CreateTeacher {
	return &createTeacherImpl{Teacher: teacherRepo}
}

// Handle creates a new teacher.
func (s *createTeacherImpl) Handle(
	ctx context.Context,
	input *CreateInput,
) (*Output, error) {
	teacher, err := model.NewTeacher(0, input.Name, input.Degree)
	if err != nil {
		return nil, err
	}

	createTeacherParams := repo.CreateTeacherParams{
		Teacher: teacher,
	}
	teacher, err = s.Teacher.Create(ctx, createTeacherParams)
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&Output{}, teacher), nil
}
