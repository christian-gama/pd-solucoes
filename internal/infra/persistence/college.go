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

type collegeImpl struct {
	db *gorm.DB
}

// NewCollege returns a new College.
func NewCollege(db *gorm.DB) repo.College {
	return &collegeImpl{
		db: db,
	}
}

// Create implements repo.College.
func (p *collegeImpl) Create(
	ctx context.Context,
	params repo.CreateCollegeParams,
) (*model.College, error) {
	db := p.db.WithContext(ctx)

	collegeSchema := convert.FromModel(&schema.College{}, &params.College)

	if err := db.
		Omit(clause.Associations).
		Create(collegeSchema).
		Error; err != nil {
		return nil, sql.Error(err, "college")
	}

	return convert.ToModel(&model.College{}, collegeSchema), nil
}

// Delete implements repo.College.
func (p *collegeImpl) Delete(ctx context.Context, params repo.DeleteCollegeParams) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.College{}).
		Error; err != nil {
		return sql.Error(err, "college")
	}

	return nil
}

// FindAll implements repo.College.
func (p *collegeImpl) FindAll(
	ctx context.Context,
	params repo.FindAllCollegeParams,
	preload ...string,
) (*queryingPort.PaginationOutput[*model.College], error) {
	db := p.db.WithContext(ctx)

	var collegeWithCount []querying.PaginationOutput[schema.College]

	if err := db.
		Model(&schema.College{}).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&collegeWithCount).
		Error; err != nil {
		return nil, sql.Error(err, "college")
	}

	pagination := &queryingPort.PaginationOutput[*model.College]{}
	for _, schema := range collegeWithCount {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.College{}, schema.Schema),
		)
		pagination.Total = int(schema.Total)
	}

	return pagination, nil
}

// FindOne implements repo.College.
func (p *collegeImpl) FindOne(
	ctx context.Context,
	params repo.FindOneCollegeParams,
	preload ...string,
) (*model.College, error) {
	db := p.db.WithContext(ctx)

	var collegeSchema schema.College

	if err := db.
		Model(&collegeSchema).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
		).
		Where("id = ?", params.ID).
		First(&collegeSchema).
		Error; err != nil {
		return nil, sql.Error(err, "college")
	}

	return convert.ToModel(&model.College{}, &collegeSchema), nil
}

// Update implements repo.College.
func (p *collegeImpl) Update(
	ctx context.Context,
	params repo.UpdateCollegeParams,
) (*model.College, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.FindOne(ctx, repo.FindOneCollegeParams{ID: params.College.ID}); err != nil {
		return nil, err
	}

	collegeSchema := convert.FromModel(&schema.College{}, &params.College)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.College.ID).
		Updates(collegeSchema).
		Error; err != nil {
		return nil, sql.Error(err, "college")
	}

	return convert.ToModel(&model.College{}, collegeSchema), nil
}
