package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllInput struct {
	Filter querying.Filter `validate:"query,filter=studentID enrollmentDate courseSubjectID"  form:"filter" faker:"-"` // nolint: revive
	Sort   querying.Sort   `validate:"query,sort=id studentID enrollmentDate courseSubjectID" form:"sort"   faker:"-"` // nolint: revive
	querying.Pagination
}
