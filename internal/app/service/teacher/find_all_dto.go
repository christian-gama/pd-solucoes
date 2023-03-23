package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllTeachersInput struct {
	Filter querying.Filter `validate:"query,filter=name degree"  form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=id name degree" form:"sort"   faker:"-"`
	querying.Pagination
}
