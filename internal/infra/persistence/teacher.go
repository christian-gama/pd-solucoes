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

type teacherImpl struct {
	db *gorm.DB
}

// NewTeacher returns a new Teacher.
func NewTeacher(db *gorm.DB) repo.Teacher {
	return &teacherImpl{
		db: db,
	}
}

// Create implements repo.Teacher.
func (p *teacherImpl) Create(
	ctx context.Context,
	params repo.CreateTeacherParams,
) (*model.Teacher, error) {
	db := p.db.WithContext(ctx)

	teacherSchema := convert.FromModel(&schema.Teacher{}, &params.Teacher)

	if err := db.
		Omit(clause.Associations).
		Create(teacherSchema).
		Error; err != nil {
		return nil, sql.Error(err, "teacher")
	}

	return convert.ToModel(&model.Teacher{}, teacherSchema), nil
}

// Delete implements repo.Teacher.
func (p *teacherImpl) Delete(ctx context.Context, params repo.DeleteTeacherParams) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.Teacher{}).
		Error; err != nil {
		return sql.Error(err, "teacher")
	}

	return nil
}

// FindAll implements repo.Teacher.
func (p *teacherImpl) FindAll(
	ctx context.Context,
	params repo.FindAllTeacherParams,
	preload ...string,
) (*queryingPort.PaginationOutput[*model.Teacher], error) {
	db := p.db.WithContext(ctx)

	var teachers []*schema.Teacher

	if err := db.
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&teachers).
		Error; err != nil {
		return nil, sql.Error(err, "teachers")
	}

	var totalCount int64
	err := db.
		Model(&schema.Teacher{}).
		Scopes(querying.FilterScope(params.Filterer)).
		Count(&totalCount).Error
	if err != nil {
		return nil, sql.Error(err, "teachers")
	}

	pagination := &queryingPort.PaginationOutput[*model.Teacher]{}
	for _, teacherSchema := range teachers {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.Teacher{}, teacherSchema),
		)
	}
	pagination.Total = int(totalCount)

	return pagination, nil
}

// FindOne implements repo.Teacher.
func (p *teacherImpl) FindOne(
	ctx context.Context,
	params repo.FindOneTeacherParams,
	preload ...string,
) (*model.Teacher, error) {
	db := p.db.WithContext(ctx)

	var teacherSchema schema.Teacher

	if err := db.
		Model(&teacherSchema).
		Scopes(sql.PreloadScope(preload)).
		Where("id = ?", params.ID).
		First(&teacherSchema).
		Error; err != nil {
		return nil, sql.Error(err, "teacher")
	}

	return convert.ToModel(&model.Teacher{}, &teacherSchema), nil
}

// Update implements repo.Teacher.
func (p *teacherImpl) Update(
	ctx context.Context,
	params repo.UpdateTeacherParams,
) (*model.Teacher, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.FindOne(ctx, repo.FindOneTeacherParams{ID: params.Teacher.ID}); err != nil {
		return nil, err
	}

	teacherSchema := convert.FromModel(&schema.Teacher{}, &params.Teacher)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.Teacher.ID).
		Updates(teacherSchema).
		Error; err != nil {
		return nil, sql.Error(err, "teacher")
	}

	return convert.ToModel(&model.Teacher{}, teacherSchema), nil
}
