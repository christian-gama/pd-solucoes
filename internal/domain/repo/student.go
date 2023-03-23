package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateStudentParams struct {
		Student *model.Student
	}

	UpdateStudentParams struct {
		Student *model.Student
	}

	DeleteStudentParams struct {
		ID uint
	}

	FindAllStudentParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindOneStudentParams struct {
		ID uint
	}
)

type Student interface {
	// Create creates a new student.
	Create(
		ctx context.Context,
		params CreateStudentParams,
	) (*model.Student, error)

	// Delete deletes a student by its ID.
	Delete(
		ctx context.Context,
		params DeleteStudentParams,
	) error

	// FindAll returns a list of students.
	FindAll(
		ctx context.Context,
		params FindAllStudentParams,
		preload ...string,
	) (*querying.PaginationOutput[*model.Student], error)

	// FindOne returns a student by its ID.
	FindOne(
		ctx context.Context,
		params FindOneStudentParams,
		preload ...string,
	) (*model.Student, error)

	// Update updates a student.
	Update(
		ctx context.Context,
		params UpdateStudentParams,
	) (*model.Student, error)
}
