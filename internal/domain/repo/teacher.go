package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateTeacherParams struct {
		Teacher *model.Teacher
	}

	UpdateTeacherParams struct {
		Teacher *model.Teacher
	}

	DeleteTeacherParams struct {
		ID uint
	}

	FindAllTeacherParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindOneTeacherParams struct {
		ID uint

		Filterer querying.Filterer
	}
)

type Teacher interface {
	// Create creates a new teacher.
	Create(
		ctx context.Context,
		params CreateTeacherParams,
	) (*model.Teacher, error)

	// Delete deletes a teacher by its ID.
	Delete(
		ctx context.Context,
		params DeleteTeacherParams,
	) error

	// FindAll returns a list of teachers.
	FindAll(
		ctx context.Context,
		params FindAllTeacherParams,
	) (*querying.PaginationOutput[*model.Teacher], error)

	// FindOne returns a teacher by its ID.
	FindOne(
		ctx context.Context,
		params FindOneTeacherParams,
	) (*model.Teacher, error)

	// Update updates a teacher.
	Update(
		ctx context.Context,
		params UpdateTeacherParams,
	) (*model.Teacher, error)
}
