package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllInput struct {
	Filter querying.Filter `validate:"query,filter=name cnpj"  form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=id name cnpj" form:"sort"   faker:"-"`
	querying.Pagination
}
