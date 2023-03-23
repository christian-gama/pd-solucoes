package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllCollegesInput struct {
	Filter querying.Filter `validate:"query,filter=name cnpj" form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=name cnpj"   form:"sort"   faker:"-"`
	querying.Pagination
}
