package service

import (
	queryingPort "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllStudentsInput struct {
	Filter querying.Filter `validate:"query,filter=name cpf" form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=name cpf"   form:"sort"   faker:"-"`
	querying.Pagination
}

type FindAllStudentsOutput = queryingPort.PaginationOutput[*FindOneStudentOutput]
