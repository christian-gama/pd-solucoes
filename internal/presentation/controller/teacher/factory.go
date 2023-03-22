package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/teacher"

func MakeCreateTeacher() CreateTeacher {
	return NewCreateTeacher(service.MakeCreateTeacher())
}

func MakeUpdateTeacher() UpdateTeacher {
	return NewUpdateTeacher(service.MakeUpdateTeacher())
}

func MakeFindOneTeacher() FindOneTeacher {
	return NewFindOneTeacher(service.MakeFindOneTeacher())
}

func MakeFindAllTeachers() FindAllTeachers {
	return NewFindAllTeachers(service.MakeFindAllTeachers())
}

func MakeDeleteTeacher() DeleteTeacher {
	return NewDeleteTeacher(service.MakeDeleteTeacher())
}
