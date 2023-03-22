package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllStudents interface {
	// Handle finds all students.
	Handle(ctx context.Context, input *FindAllStudentsInput) (*FindAllStudentsOutput, error)
}

type findAllStudentsImpl struct {
	repo.Student
}

// NewFindAllStudents returns a FindAllStudents.
func NewFindAllStudents(studentRepo repo.Student) FindAllStudents {
	return &findAllStudentsImpl{Student: studentRepo}
}

// Handle findAlls a new student.
func (s *findAllStudentsImpl) Handle(
	ctx context.Context,
	input *FindAllStudentsInput,
) (*FindAllStudentsOutput, error) {
	findAllStudentParams := repo.FindAllStudentParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	student, err := s.Student.FindAll(ctx, findAllStudentParams)
	if err != nil {
		return nil, err
	}

	result := make([]*FindOneStudentOutput, 0, len(student.Results))
	for _, c := range student.Results {
		result = append(result, &FindOneStudentOutput{
			ID:   c.ID,
			Name: c.Name,
			Cpf:  c.Cpf,
		})
	}

	output := &FindAllStudentsOutput{
		Total:   student.Total,
		Results: result,
	}

	return output, nil
}
