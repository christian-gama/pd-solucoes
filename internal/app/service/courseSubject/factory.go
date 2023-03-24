package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateCourseSubject() CreateCourseSubject {
	return NewCreateCourseSubject(persistence.MakeCourseSubject())
}

func MakeUpdateCourseSubject() UpdateCourseSubject {
	return NewUpdateCourseSubject(persistence.MakeCourseSubject())
}

func MakeFindOneCourseSubject() FindOneCourseSubject {
	return NewFindOneCourseSubject(persistence.MakeCourseSubject())
}

func MakeFindAllCourseSubjects() FindAllCourseSubjects {
	return NewFindAllCourseSubjects(persistence.MakeCourseSubject())
}

func MakeDeleteCourseSubject() DeleteCourseSubject {
	return NewDeleteCourseSubject(persistence.MakeCourseSubject())
}
