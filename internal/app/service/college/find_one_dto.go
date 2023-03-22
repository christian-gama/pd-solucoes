package service

import "github.com/christian-gama/pd-solucoes/internal/infra/querying"

type FindOneCollegeInput struct {
	ID     uint            `validate:"required"               uri:"id"`
	Filter querying.Filter `validate:"query,filter=name cnpj"          form:"filter" faker:"-"`
}

type FindOneCollegeOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cnpj string `json:"cnpj"`
}
