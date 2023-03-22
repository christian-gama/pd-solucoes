package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateCourseSubjectParams struct {
		CourseSubject *model.CourseSubject
	}

	UpdateCourseSubjectParams struct {
		CourseSubject *model.CourseSubject
	}

	DeleteCourseSubjectParams struct {
		ID uint
	}

	FindAllCourseSubjectParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindOneCourseSubjectParams struct {
		ID uint
	}
)

type CourseSubject interface {
	// Create creates a new courseSubject.
	Create(
		ctx context.Context,
		params CreateCourseSubjectParams,
	) (*model.CourseSubject, error)

	// Delete deletes a courseSubject by its ID.
	Delete(
		ctx context.Context,
		params DeleteCourseSubjectParams,
	) error

	// FindAll returns a list of courseSubjects.
	FindAll(
		ctx context.Context,
		params FindAllCourseSubjectParams,
		preload ...string,
	) (*querying.PaginationOutput[*model.CourseSubject], error)

	// FindOne returns a courseSubject by its ID.
	FindOne(
		ctx context.Context,
		params FindOneCourseSubjectParams,
		preload ...string,
	) (*model.CourseSubject, error)

	// Update updates a courseSubject.
	Update(
		ctx context.Context,
		params UpdateCourseSubjectParams,
	) (*model.CourseSubject, error)
}
