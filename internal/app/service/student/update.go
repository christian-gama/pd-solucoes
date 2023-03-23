package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type UpdateStudent interface {
	// Handle updates a student.
	Handle(ctx context.Context, input *UpdateStudentInput) (*model.Student, error)
}

type updateStudentImpl struct {
	repo.Student
}

// NewUpdateStudent returns a UpdateStudent.
func NewUpdateStudent(studentRepo repo.Student) UpdateStudent {
	return &updateStudentImpl{Student: studentRepo}
}

// Handle updates a student.
func (s *updateStudentImpl) Handle(
	ctx context.Context,
	input *UpdateStudentInput,
) (*model.Student, error) {
	student, err := model.NewStudent(input.ID, input.Name, input.Cpf)
	if err != nil {
		return nil, err
	}

	updateStudentParams := repo.UpdateStudentParams{
		Student: student,
	}
	student, err = s.Student.Update(ctx, updateStudentParams)
	if err != nil {
		return nil, err
	}

	return student, nil
}
