package dto

import (
	queryingPort "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindOneCollegeInput struct {
	ID     uint            `validate:"required"               uri:"id"`
	Filter querying.Filter `validate:"query,filter=name cnpj"          form:"filter" faker:"-"`
}

type FindOneCollegeOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}

type CreateCollegeInput struct {
	Name string `json:"name" validate:"required,max=100,min=2" faker:"len=50"`
	Cnpj string `json:"cnpj" validate:"required,cnpj"          faker:"cnpj"`
}

type CreateCollegeOutput = FindOneCollegeOutput

type UpdateCollegeInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
	CreateCollegeInput
}

type UpdateCollegeOutput = FindOneCollegeOutput

type FindAllCollegesInput struct {
	Filter querying.Filter `validate:"query,filter=name cnpj" form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=name cnpj"   form:"sort"   faker:"-"`
	querying.Pagination
}

type FindAllCollegesOutput = queryingPort.PaginationOutput[*FindOneCollegeOutput]

type DeleteCollegeInput struct {
	ID uint `uri:"id" validate:"required" faker:"uint"`
}
