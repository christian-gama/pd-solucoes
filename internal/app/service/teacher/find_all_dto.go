package service

import (
	queryingPort "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllTeachersInput struct {
	Filter querying.Filter `validate:"query,filter=name degree" form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=name degree"   form:"sort"   faker:"-"`
	querying.Pagination
}

type FindAllTeachersOutput = queryingPort.PaginationOutput[*FindOneTeacherOutput]
