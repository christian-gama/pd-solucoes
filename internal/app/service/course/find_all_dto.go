package service

import (
	queryingPort "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
)

type FindAllCoursesInput struct {
	Filter querying.Filter `validate:"query,filter=name collegeID" form:"filter" faker:"-"`
	Sort   querying.Sort   `validate:"query,sort=name collegeID"   form:"sort"   faker:"-"`
	querying.Pagination
}

type FindAllCoursesOutput = queryingPort.PaginationOutput[*FindOneCourseOutput]
