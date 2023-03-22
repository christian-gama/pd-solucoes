package persistence

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	queryingPort "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/internal/infra/convert"
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence/schema"
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type courseEnrollmentImpl struct {
	db *gorm.DB
}

// NewCourseEnrollment returns a new CourseEnrollment.
func NewCourseEnrollment(db *gorm.DB) repo.CourseEnrollment {
	return &courseEnrollmentImpl{
		db: db,
	}
}

// Create implements repo.CourseEnrollment.
func (p *courseEnrollmentImpl) Create(
	ctx context.Context,
	params repo.CreateCourseEnrollmentParams,
) (*model.CourseEnrollment, error) {
	db := p.db.WithContext(ctx)

	courseEnrollmentSchema := convert.FromModel(&schema.CourseEnrollment{}, &params.CourseEnrollment)

	if err := db.
		Omit(clause.Associations).
		Create(courseEnrollmentSchema).
		Error; err != nil {
		return nil, sql.Error(err, "courseEnrollment")
	}

	return convert.ToModel(&model.CourseEnrollment{}, courseEnrollmentSchema), nil
}

// Delete implements repo.CourseEnrollment.
func (p *courseEnrollmentImpl) Delete(
	ctx context.Context,
	params repo.DeleteCourseEnrollmentParams,
) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.CourseEnrollment{}).
		Error; err != nil {
		return sql.Error(err, "courseEnrollment")
	}

	return nil
}

// FindAll implements repo.CourseEnrollment.
func (p *courseEnrollmentImpl) FindAll(
	ctx context.Context,
	params repo.FindAllCourseEnrollmentParams,
	preload ...string,
) (*queryingPort.PaginationOutput[*model.CourseEnrollment], error) {
	db := p.db.WithContext(ctx)

	var courseEnrollmentWithCount []querying.PaginationOutput[schema.CourseEnrollment]

	if err := db.
		Model(&schema.CourseEnrollment{}).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&courseEnrollmentWithCount).
		Error; err != nil {
		return nil, sql.Error(err, "courseEnrollment")
	}

	pagination := &queryingPort.PaginationOutput[*model.CourseEnrollment]{}
	for _, schema := range courseEnrollmentWithCount {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.CourseEnrollment{}, schema.Schema),
		)
		pagination.Total = int(schema.Total)
	}

	return pagination, nil
}

// FindOne implements repo.CourseEnrollment.
func (p *courseEnrollmentImpl) FindOne(
	ctx context.Context,
	params repo.FindOneCourseEnrollmentParams,
	preload ...string,
) (*model.CourseEnrollment, error) {
	db := p.db.WithContext(ctx)

	var courseEnrollmentSchema schema.CourseEnrollment

	if err := db.
		Model(&courseEnrollmentSchema).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
		).
		Where("id = ?", params.ID).
		First(&courseEnrollmentSchema).
		Error; err != nil {
		return nil, sql.Error(err, "courseEnrollment")
	}

	return convert.ToModel(&model.CourseEnrollment{}, &courseEnrollmentSchema), nil
}

// Update implements repo.CourseEnrollment.
func (p *courseEnrollmentImpl) Update(
	ctx context.Context,
	params repo.UpdateCourseEnrollmentParams,
) (*model.CourseEnrollment, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.
		FindOne(ctx, repo.FindOneCourseEnrollmentParams{ID: params.CourseEnrollment.ID}); err != nil {
		return nil, err
	}

	courseEnrollmentSchema := convert.FromModel(&schema.CourseEnrollment{}, &params.CourseEnrollment)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.CourseEnrollment.ID).
		Updates(courseEnrollmentSchema).
		Error; err != nil {
		return nil, sql.Error(err, "courseEnrollment")
	}

	return convert.ToModel(&model.CourseEnrollment{}, courseEnrollmentSchema), nil
}
