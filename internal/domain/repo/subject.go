package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateSubjectParams struct {
		Subject *model.Subject
	}

	UpdateSubjectParams struct {
		Subject *model.Subject
	}

	DeleteSubjectParams struct {
		ID uint
	}

	FindAllSubjectParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindOneSubjectParams struct {
		ID uint
	}
)

type Subject interface {
	// Create creates a new subject.
	Create(
		ctx context.Context,
		params CreateSubjectParams,
	) (*model.Subject, error)

	// Delete deletes a subject by its ID.
	Delete(
		ctx context.Context,
		params DeleteSubjectParams,
	) error

	// FindAll returns a list of subjects.
	FindAll(
		ctx context.Context,
		params FindAllSubjectParams,
		preload ...string,
	) (*querying.PaginationOutput[*model.Subject], error)

	// FindOne returns a subject by its ID.
	FindOne(
		ctx context.Context,
		params FindOneSubjectParams,
		preload ...string,
	) (*model.Subject, error)

	// Update updates a subject.
	Update(
		ctx context.Context,
		params UpdateSubjectParams,
	) (*model.Subject, error)
}
