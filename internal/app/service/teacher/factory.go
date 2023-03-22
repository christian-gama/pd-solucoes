package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateTeacher() CreateTeacher {
	return NewCreateTeacher(persistence.MakeTeacher())
}

func MakeUpdateTeacher() UpdateTeacher {
	return NewUpdateTeacher(persistence.MakeTeacher())
}

func MakeFindOneTeacher() FindOneTeacher {
	return NewFindOneTeacher(persistence.MakeTeacher())
}

func MakeFindAllTeachers() FindAllTeachers {
	return NewFindAllTeachers(persistence.MakeTeacher())
}

func MakeDeleteTeacher() DeleteTeacher {
	return NewDeleteTeacher(persistence.MakeTeacher())
}
