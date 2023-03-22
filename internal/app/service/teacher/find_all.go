package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllTeachers interface {
	// Handle finds all teachers.
	Handle(ctx context.Context, input *FindAllTeachersInput) (*FindAllTeachersOutput, error)
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
) (*FindAllTeachersOutput, error) {
	findAllTeacherParams := repo.FindAllTeacherParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	teacher, err := s.Teacher.FindAll(ctx, findAllTeacherParams)
	if err != nil {
		return nil, err
	}

	result := make([]*FindOneTeacherOutput, 0, len(teacher.Results))
	for _, c := range teacher.Results {
		result = append(result, &FindOneTeacherOutput{
			ID:     c.ID,
			Name:   c.Name,
			Degree: c.Degree,
		})
	}

	output := &FindAllTeachersOutput{
		Total:   teacher.Total,
		Results: result,
	}

	return output, nil
}
