package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllStudentsInput struct {
	Filter querying.Filter `validate:"query,filter=name cpf"  form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=id name cpf" form:"sort"   faker:"-"`
	querying.Pagination
}
