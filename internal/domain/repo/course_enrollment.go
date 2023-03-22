package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateCourseEnrollmentParams struct {
		CourseEnrollment *model.CourseEnrollment
	}

	UpdateCourseEnrollmentParams struct {
		CourseEnrollment *model.CourseEnrollment
	}

	DeleteCourseEnrollmentParams struct {
		ID uint
	}

	FindAllCourseEnrollmentParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindOneCourseEnrollmentParams struct {
		ID uint
	}
)

type CourseEnrollment interface {
	// Create creates a new courseEnrollment.
	Create(
		ctx context.Context,
		params CreateCourseEnrollmentParams,
	) (*model.CourseEnrollment, error)

	// Delete deletes a courseEnrollment by its ID.
	Delete(
		ctx context.Context,
		params DeleteCourseEnrollmentParams,
	) error

	// FindAll returns a list of courseEnrollments.
	FindAll(
		ctx context.Context,
		params FindAllCourseEnrollmentParams,
		preload ...string,
	) (*querying.PaginationOutput[*model.CourseEnrollment], error)

	// FindOne returns a courseEnrollment by its ID.
	FindOne(
		ctx context.Context,
		params FindOneCourseEnrollmentParams,
		preload ...string,
	) (*model.CourseEnrollment, error)

	// Update updates a courseEnrollment.
	Update(
		ctx context.Context,
		params UpdateCourseEnrollmentParams,
	) (*model.CourseEnrollment, error)
}
