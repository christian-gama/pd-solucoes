package querying

import (
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"gorm.io/gorm"
)

// PaginationOutput is a struct used to return paginated results directly from the database.
type PaginationOutput[Schema any] struct {
	Total  int64  `gorm:"column:total"`
	Schema Schema `gorm:"embedded"`
}

// Pagination is a struct used to paginate queries.
type Pagination struct {
	Page  int `form:"page"  validate:"omitempty,min=1"`
	Limit int `form:"limit" validate:"omitempty,max=100,min=1"`
}

// GetLimit implements querying.Paginator.
func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Limit > 100 {
		p.Limit = 100
	}

	return p.Limit
}

// GetOffset implements querying.Paginator.
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetPage implements querying.Paginator.
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}

	return p.Page
}

// SetLimit implements querying.Paginator.
func (p *Pagination) SetLimit(limit int) querying.Paginator {
	p.Limit = limit
	return p
}

// SetPage implements querying.Paginator.
func (p *Pagination) SetPage(page int) querying.Paginator {
	p.Page = page
	return p
}

// PaginationScope is a function that can be used as a GORM scope to paginate queries.
func PaginationScope(paginator querying.Paginator) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paginator == nil {
			return db
		}

		return db.
			Offset(paginator.GetOffset()).
			Limit(paginator.GetLimit()).
			Select("*, COUNT(*) OVER() AS total")
	}
}
