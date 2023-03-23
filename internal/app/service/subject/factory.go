package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateSubject() CreateSubject {
	return NewCreateSubject(persistence.MakeSubject())
}

func MakeUpdateSubject() UpdateSubject {
	return NewUpdateSubject(persistence.MakeSubject())
}

func MakeFindOneSubject() FindOneSubject {
	return NewFindOneSubject(persistence.MakeSubject())
}

func MakeFindAllSubjects() FindAllSubjects {
	return NewFindAllSubjects(persistence.MakeSubject())
}

func MakeDeleteSubject() DeleteSubject {
	return NewDeleteSubject(persistence.MakeSubject())
}
