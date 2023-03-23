package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllTeachers interface {
	// Handle finds all teachers.
	Handle(
		ctx context.Context,
		input *FindAllTeachersInput,
	) (*querying.PaginationOutput[*model.Teacher], error)
}

type findAllTeachersImpl struct {
	repo.Teacher
}

// NewFindAllTeachers returns a FindAllTeachers.
func NewFindAllTeachers(teacherRepo repo.Teacher) FindAllTeachers {
	return &findAllTeachersImpl{Teacher: teacherRepo}
}

// Handle findAlls a new teacher.
func (s *findAllTeachersImpl) Handle(
	ctx context.Context,
	input *FindAllTeachersInput,
) (*querying.PaginationOutput[*model.Teacher], error) {
	findAllTeacherParams := repo.FindAllTeacherParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	paginationOutput, err := s.Teacher.FindAll(ctx, findAllTeacherParams, "subjects")
	if err != nil {
		return nil, err
	}

	return paginationOutput, nil
}
