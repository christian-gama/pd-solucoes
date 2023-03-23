package service

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

type FindAllColleges interface {
	// Handle finds all colleges.
	Handle(ctx context.Context, input *FindAllCollegesInput) (*FindAllCollegesOutput, error)
}

type findAllCollegesImpl struct {
	repo.College
}

// NewFindAllColleges returns a FindAllColleges.
func NewFindAllColleges(collegeRepo repo.College) FindAllColleges {
	return &findAllCollegesImpl{College: collegeRepo}
}

// Handle findAlls a new college.
func (s *findAllCollegesImpl) Handle(
	ctx context.Context,
	input *FindAllCollegesInput,
) (*FindAllCollegesOutput, error) {
	findAllCollegeParams := repo.FindAllCollegeParams{
		Paginator: &input.Pagination,
		Filterer:  input.Filter,
		Sorter:    input.Sort,
	}
	college, err := s.College.FindAll(ctx, findAllCollegeParams, "courses")
	if err != nil {
		return nil, err
	}

	result := make([]*FindOneCollegeOutput, len(college.Results))
	for i, c := range college.Results {
		coursesOutput := make([]*findOneCollegeCourseOutput, len(c.Courses))
		for j, course := range c.Courses {
			coursesOutput[j] = &findOneCollegeCourseOutput{
				ID:   course.ID,
				Name: course.Name,
			}
		}

		result[i] = &FindOneCollegeOutput{
			ID:      c.ID,
			Name:    c.Name,
			Cnpj:    c.Cnpj,
			Courses: coursesOutput,
		}
	}

	output := &FindAllCollegesOutput{
		Total:   college.Total,
		Results: result,
	}

	return output, nil
}
