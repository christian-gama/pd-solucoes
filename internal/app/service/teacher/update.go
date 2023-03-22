package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type UpdateTeacher interface {
	// Handle updates a teacher.
	Handle(ctx context.Context, input *UpdateTeacherInput) (*UpdateTeacherOutput, error)
}

type updateTeacherImpl struct {
	repo.Teacher
}

// NewUpdateTeacher returns a UpdateTeacher.
func NewUpdateTeacher(teacherRepo repo.Teacher) UpdateTeacher {
	return &updateTeacherImpl{Teacher: teacherRepo}
}

// Handle updates a teacher.
func (s *updateTeacherImpl) Handle(
	ctx context.Context,
	input *UpdateTeacherInput,
) (*UpdateTeacherOutput, error) {
	teacher, err := model.NewTeacher(input.ID, input.Name, input.Degree)
	if err != nil {
		return nil, err
	}

	updateTeacherParams := repo.UpdateTeacherParams{
		Teacher: teacher,
	}
	teacher, err = s.Teacher.Update(ctx, updateTeacherParams)
	if err != nil {
		return nil, err
	}

	output := &UpdateTeacherOutput{
		ID:     teacher.ID,
		Name:   teacher.Name,
		Degree: teacher.Degree,
	}

	return output, nil
}
