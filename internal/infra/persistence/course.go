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

type courseImpl struct {
	db *gorm.DB
}

// NewCourse returns a new Course.
func NewCourse(db *gorm.DB) repo.Course {
	return &courseImpl{
		db: db,
	}
}

// Create implements repo.Course.
func (p *courseImpl) Create(
	ctx context.Context,
	params repo.CreateCourseParams,
) (*model.Course, error) {
	db := p.db.WithContext(ctx)

	courseSchema := convert.FromModel(&schema.Course{}, &params.Course)

	if err := db.
		Omit(clause.Associations).
		Create(courseSchema).
		Error; err != nil {
		return nil, sql.Error(err, "course")
	}

	return convert.ToModel(&model.Course{}, courseSchema), nil
}

// Delete implements repo.Course.
func (p *courseImpl) Delete(ctx context.Context, params repo.DeleteCourseParams) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.Course{}).
		Error; err != nil {
		return sql.Error(err, "course")
	}

	return nil
}

// FindAll implements repo.Course.
func (p *courseImpl) FindAll(
	ctx context.Context,
	params repo.FindAllCourseParams,
	preload ...string,
) (*queryingPort.PaginationOutput[*model.Course], error) {
	db := p.db.WithContext(ctx)

	var courseWithCount []querying.PaginationOutput[schema.Course]

	if err := db.
		Model(&schema.Course{}).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&courseWithCount).
		Error; err != nil {
		return nil, sql.Error(err, "course")
	}

	pagination := &queryingPort.PaginationOutput[*model.Course]{}
	for _, schema := range courseWithCount {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.Course{}, schema.Schema),
		)
		pagination.Total = int(schema.Total)
	}

	return pagination, nil
}

// FindOne implements repo.Course.
func (p *courseImpl) FindOne(
	ctx context.Context,
	params repo.FindOneCourseParams,
	preload ...string,
) (*model.Course, error) {
	db := p.db.WithContext(ctx)

	var courseSchema schema.Course

	if err := db.
		Model(&courseSchema).
		Scopes(sql.PreloadScope(preload)).
		Where("id = ?", params.ID).
		First(&courseSchema).
		Error; err != nil {
		return nil, sql.Error(err, "course")
	}

	return convert.ToModel(&model.Course{}, &courseSchema), nil
}

// Update implements repo.Course.
func (p *courseImpl) Update(
	ctx context.Context,
	params repo.UpdateCourseParams,
) (*model.Course, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.FindOne(ctx, repo.FindOneCourseParams{ID: params.Course.ID}); err != nil {
		return nil, err
	}

	courseSchema := convert.FromModel(&schema.Course{}, &params.Course)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.Course.ID).
		Updates(courseSchema).
		Error; err != nil {
		return nil, sql.Error(err, "course")
	}

	return convert.ToModel(&model.Course{}, courseSchema), nil
}
