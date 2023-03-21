package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateCollegeParams struct {
		College *model.College
	}

	UpdateCollegeParams struct {
		College *model.College
	}

	DeleteCollegeParams struct {
		ID uint
	}

	FindAllCollegeParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindCollegeByIDParams struct {
		ID uint

		Filterer querying.Filterer
	}
)

type College interface {
	// Create creates a new college.
	Create(
		ctx context.Context,
		params *CreateCollegeParams,
	) (*model.College, error)

	// Delete deletes a college by its ID.
	Delete(
		ctx context.Context,
		params *DeleteCollegeParams,
	) error

	// FindAll returns a list of colleges.
	FindAll(
		ctx context.Context,
		params *FindAllCollegeParams,
		preload ...string,
	) (*querying.PaginationOutput[*model.College], error)

	// FindByID returns a college by its ID.
	FindByID(
		ctx context.Context,
		params *FindCollegeByIDParams,
		preload ...string,
	) (*model.College, error)

	// Update updates a college.
	Update(
		ctx context.Context,
		params *UpdateCollegeParams,
	) (*model.College, error)
}
