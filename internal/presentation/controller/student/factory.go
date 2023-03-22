package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/student"

func MakeCreateStudent() CreateStudent {
	return NewCreateStudent(service.MakeCreateStudent())
}

func MakeUpdateStudent() UpdateStudent {
	return NewUpdateStudent(service.MakeUpdateStudent())
}

func MakeFindOneStudent() FindOneStudent {
	return NewFindOneStudent(service.MakeFindOneStudent())
}

func MakeFindAllStudents() FindAllStudents {
	return NewFindAllStudents(service.MakeFindAllStudents())
}

func MakeDeleteStudent() DeleteStudent {
	return NewDeleteStudent(service.MakeDeleteStudent())
}
