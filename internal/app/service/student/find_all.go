package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllStudents interface {
	// Handle finds all students.
	Handle(
		ctx context.Context,
		input *FindAllStudentsInput,
	) (*querying.PaginationOutput[*model.Student], error)
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
) (*querying.PaginationOutput[*model.Student], error) {
	findAllStudentParams := repo.FindAllStudentParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.Student.FindAll(ctx, findAllStudentParams)
	if err != nil {
		return nil, err
	}

	return paginationOutput, nil
}
