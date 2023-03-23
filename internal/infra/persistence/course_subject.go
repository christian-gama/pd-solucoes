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

type courseSubjectImpl struct {
	db *gorm.DB
}

// NewCourseSubject returns a new CourseSubject.
func NewCourseSubject(db *gorm.DB) repo.CourseSubject {
	return &courseSubjectImpl{
		db: db,
	}
}

// Create implements repo.CourseSubject.
func (p *courseSubjectImpl) Create(
	ctx context.Context,
	params repo.CreateCourseSubjectParams,
) (*model.CourseSubject, error) {
	db := p.db.WithContext(ctx)

	courseSubjectSchema := convert.FromModel(&schema.CourseSubject{}, &params.CourseSubject)

	if err := db.
		Omit(clause.Associations).
		Create(courseSubjectSchema).
		Error; err != nil {
		return nil, sql.Error(err, "courseSubject")
	}

	return convert.ToModel(&model.CourseSubject{}, courseSubjectSchema), nil
}

// Delete implements repo.CourseSubject.
func (p *courseSubjectImpl) Delete(
	ctx context.Context,
	params repo.DeleteCourseSubjectParams,
) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.CourseSubject{}).
		Error; err != nil {
		return sql.Error(err, "courseSubject")
	}

	return nil
}

// FindAll implements repo.CourseSubject.
func (p *courseSubjectImpl) FindAll(
	ctx context.Context,
	params repo.FindAllCourseSubjectParams,
	preload ...string,
) (*queryingPort.PaginationOutput[*model.CourseSubject], error) {
	db := p.db.WithContext(ctx)

	var courseSubjectWithCount []querying.PaginationOutput[schema.CourseSubject]

	if err := db.
		Model(&schema.CourseSubject{}).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&courseSubjectWithCount).
		Error; err != nil {
		return nil, sql.Error(err, "courseSubject")
	}

	pagination := &queryingPort.PaginationOutput[*model.CourseSubject]{}
	for _, schema := range courseSubjectWithCount {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.CourseSubject{}, schema.Schema),
		)
		pagination.Total = int(schema.Total)
	}

	return pagination, nil
}

// FindOne implements repo.CourseSubject.
func (p *courseSubjectImpl) FindOne(
	ctx context.Context,
	params repo.FindOneCourseSubjectParams,
	preload ...string,
) (*model.CourseSubject, error) {
	db := p.db.WithContext(ctx)

	var courseSubjectSchema schema.CourseSubject

	if err := db.
		Model(&courseSubjectSchema).
		Scopes(sql.PreloadScope(preload)).
		Where("id = ?", params.ID).
		First(&courseSubjectSchema).
		Error; err != nil {
		return nil, sql.Error(err, "courseSubject")
	}

	return convert.ToModel(&model.CourseSubject{}, &courseSubjectSchema), nil
}

// Update implements repo.CourseSubject.
func (p *courseSubjectImpl) Update(
	ctx context.Context,
	params repo.UpdateCourseSubjectParams,
) (*model.CourseSubject, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.FindOne(ctx, repo.FindOneCourseSubjectParams{ID: params.CourseSubject.ID}); err != nil {
		return nil, err
	}

	courseSubjectSchema := convert.FromModel(&schema.CourseSubject{}, &params.CourseSubject)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.CourseSubject.ID).
		Updates(courseSubjectSchema).
		Error; err != nil {
		return nil, sql.Error(err, "courseSubject")
	}

	return convert.ToModel(&model.CourseSubject{}, courseSubjectSchema), nil
}
