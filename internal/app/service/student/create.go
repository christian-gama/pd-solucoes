package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type CreateStudent interface {
	// Handle creates a new student.
	Handle(ctx context.Context, input *CreateStudentInput) (*CreateStudentOutput, error)
}

type createStudentImpl struct {
	repo.Student
}

// NewCreateStudent returns a CreateStudent.
func NewCreateStudent(studentRepo repo.Student) CreateStudent {
	return &createStudentImpl{Student: studentRepo}
}

// Handle creates a new student.
func (s *createStudentImpl) Handle(
	ctx context.Context,
	input *CreateStudentInput,
) (*CreateStudentOutput, error) {
	student, err := model.NewStudent(0, input.Name, input.Cpf)
	if err != nil {
		return nil, err
	}

	createStudentParams := repo.CreateStudentParams{
		Student: student,
	}
	student, err = s.Student.Create(ctx, createStudentParams)
	if err != nil {
		return nil, err
	}

	output := &CreateStudentOutput{
		ID:   student.ID,
		Name: student.Name,
		Cpf:  student.Cpf,
	}

	return output, nil
}
