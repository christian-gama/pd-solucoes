package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateStudent() CreateStudent {
	return NewCreateStudent(persistence.MakeStudent())
}

func MakeUpdateStudent() UpdateStudent {
	return NewUpdateStudent(persistence.MakeStudent())
}

func MakeFindOneStudent() FindOneStudent {
	return NewFindOneStudent(persistence.MakeStudent())
}

func MakeFindAllStudents() FindAllStudents {
	return NewFindAllStudents(persistence.MakeStudent())
}

func MakeDeleteStudent() DeleteStudent {
	return NewDeleteStudent(persistence.MakeStudent())
}
