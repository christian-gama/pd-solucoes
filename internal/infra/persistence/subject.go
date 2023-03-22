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

type subjectImpl struct {
	db *gorm.DB
}

// NewSubject returns a new Subject.
func NewSubject(db *gorm.DB) repo.Subject {
	return &subjectImpl{
		db: db,
	}
}

// Create implements repo.Subject.
func (p *subjectImpl) Create(
	ctx context.Context,
	params repo.CreateSubjectParams,
) (*model.Subject, error) {
	db := p.db.WithContext(ctx)

	subjectSchema := convert.FromModel(&schema.Subject{}, &params.Subject)

	if err := db.
		Omit(clause.Associations).
		Create(subjectSchema).
		Error; err != nil {
		return nil, sql.Error(err, "subject")
	}

	return convert.ToModel(&model.Subject{}, subjectSchema), nil
}

// Delete implements repo.Subject.
func (p *subjectImpl) Delete(ctx context.Context, params repo.DeleteSubjectParams) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.Subject{}).
		Error; err != nil {
		return sql.Error(err, "subject")
	}

	return nil
}

// FindAll implements repo.Subject.
func (p *subjectImpl) FindAll(
	ctx context.Context,
	params repo.FindAllSubjectParams,
	preload ...string,
) (*queryingPort.PaginationOutput[*model.Subject], error) {
	db := p.db.WithContext(ctx)

	var subjectWithCount []querying.PaginationOutput[schema.Subject]

	if err := db.
		Model(&schema.Subject{}).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&subjectWithCount).
		Error; err != nil {
		return nil, sql.Error(err, "subject")
	}

	pagination := &queryingPort.PaginationOutput[*model.Subject]{}
	for _, schema := range subjectWithCount {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.Subject{}, schema.Schema),
		)
		pagination.Total = int(schema.Total)
	}

	return pagination, nil
}

// FindOne implements repo.Subject.
func (p *subjectImpl) FindOne(
	ctx context.Context,
	params repo.FindOneSubjectParams,
	preload ...string,
) (*model.Subject, error) {
	db := p.db.WithContext(ctx)

	var subjectSchema schema.Subject

	if err := db.
		Model(&subjectSchema).
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(params.Filterer),
		).
		Where("id = ?", params.ID).
		First(&subjectSchema).
		Error; err != nil {
		return nil, sql.Error(err, "subject")
	}

	return convert.ToModel(&model.Subject{}, &subjectSchema), nil
}

// Update implements repo.Subject.
func (p *subjectImpl) Update(
	ctx context.Context,
	params repo.UpdateSubjectParams,
) (*model.Subject, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.FindOne(ctx, repo.FindOneSubjectParams{ID: params.Subject.ID}); err != nil {
		return nil, err
	}

	subjectSchema := convert.FromModel(&schema.Subject{}, &params.Subject)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.Subject.ID).
		Updates(subjectSchema).
		Error; err != nil {
		return nil, sql.Error(err, "subject")
	}

	return convert.ToModel(&model.Subject{}, subjectSchema), nil
}
