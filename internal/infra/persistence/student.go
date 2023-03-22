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

type studentImpl struct {
	db *gorm.DB
}

// NewStudent returns a new Student.
func NewStudent(db *gorm.DB) repo.Student {
	return &studentImpl{
		db: db,
	}
}

// Create implements repo.Student.
func (p *studentImpl) Create(
	ctx context.Context,
	params repo.CreateStudentParams,
) (*model.Student, error) {
	db := p.db.WithContext(ctx)

	studentSchema := convert.FromModel(&schema.Student{}, &params.Student)

	if err := db.
		Omit(clause.Associations).
		Create(studentSchema).
		Error; err != nil {
		return nil, sql.Error(err, "student")
	}

	return convert.ToModel(&model.Student{}, studentSchema), nil
}

// Delete implements repo.Student.
func (p *studentImpl) Delete(ctx context.Context, params repo.DeleteStudentParams) error {
	db := p.db.WithContext(ctx)

	if err := db.
		Where("id = ?", params.ID).
		Delete(&schema.Student{}).
		Error; err != nil {
		return sql.Error(err, "student")
	}

	return nil
}

// FindAll implements repo.Student.
func (p *studentImpl) FindAll(
	ctx context.Context,
	params repo.FindAllStudentParams,
) (*queryingPort.PaginationOutput[*model.Student], error) {
	db := p.db.WithContext(ctx)

	var studentWithCount []querying.PaginationOutput[schema.Student]

	if err := db.
		Model(&schema.Student{}).
		Scopes(
			querying.FilterScope(params.Filterer),
			querying.PaginationScope(params.Paginator),
			querying.SortScope(params.Sorter),
		).
		Find(&studentWithCount).
		Error; err != nil {
		return nil, sql.Error(err, "student")
	}

	pagination := &queryingPort.PaginationOutput[*model.Student]{}
	for _, schema := range studentWithCount {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model.Student{}, schema.Schema),
		)
		pagination.Total = int(schema.Total)
	}

	return pagination, nil
}

// FindOne implements repo.Student.
func (p *studentImpl) FindOne(
	ctx context.Context,
	params repo.FindOneStudentParams,
) (*model.Student, error) {
	db := p.db.WithContext(ctx)

	var studentSchema schema.Student

	if err := db.
		Model(&studentSchema).
		Scopes(
			querying.FilterScope(params.Filterer),
		).
		Where("id = ?", params.ID).
		First(&studentSchema).
		Error; err != nil {
		return nil, sql.Error(err, "student")
	}

	return convert.ToModel(&model.Student{}, &studentSchema), nil
}

// Update implements repo.Student.
func (p *studentImpl) Update(
	ctx context.Context,
	params repo.UpdateStudentParams,
) (*model.Student, error) {
	db := p.db.WithContext(ctx)

	if _, err := p.FindOne(ctx, repo.FindOneStudentParams{ID: params.Student.ID}); err != nil {
		return nil, err
	}

	studentSchema := convert.FromModel(&schema.Student{}, &params.Student)

	if err := db.
		Omit(clause.Associations).
		Where("id = ?", params.Student.ID).
		Updates(studentSchema).
		Error; err != nil {
		return nil, sql.Error(err, "student")
	}

	return convert.ToModel(&model.Student{}, studentSchema), nil
}
