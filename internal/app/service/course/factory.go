package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateCourse() CreateCourse {
	return NewCreateCourse(persistence.MakeCourse())
}

func MakeUpdateCourse() UpdateCourse {
	return NewUpdateCourse(persistence.MakeCourse())
}

func MakeFindOneCourse() FindOneCourse {
	return NewFindOneCourse(persistence.MakeCourse())
}

func MakeFindAllCourses() FindAllCourses {
	return NewFindAllCourses(persistence.MakeCourse())
}

func MakeDeleteCourse() DeleteCourse {
	return NewDeleteCourse(persistence.MakeCourse())
}
