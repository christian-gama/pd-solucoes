package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllCourseSubjectsInput struct {
	Filter querying.Filter `validate:"query,filter=name courseID subjectID"  form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=id name courseID subjectID" form:"sort"   faker:"-"`
	querying.Pagination
}
