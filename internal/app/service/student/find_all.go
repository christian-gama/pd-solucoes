package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/pkg/copy"
)

type FindAllStudents interface {
	// Handle finds all students.
	Handle(
		ctx context.Context,
		input *FindAllStudentsInput,
	) (*querying.PaginationOutput[*Output], error)
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
) (*querying.PaginationOutput[*Output], error) {
	findAllStudentParams := repo.FindAllStudentParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.Student.FindAll(ctx, findAllStudentParams, "courseSubjects")
	if err != nil {
		return nil, err
	}

	return copy.MustCopy(&querying.PaginationOutput[*Output]{}, paginationOutput), nil
}
