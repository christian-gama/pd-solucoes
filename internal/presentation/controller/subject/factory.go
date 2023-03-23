package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"

func MakeCreateSubject() CreateSubject {
	return NewCreateSubject(service.MakeCreateSubject())
}

func MakeUpdateSubject() UpdateSubject {
	return NewUpdateSubject(service.MakeUpdateSubject())
}

func MakeFindOneSubject() FindOneSubject {
	return NewFindOneSubject(service.MakeFindOneSubject())
}

func MakeFindAllSubjects() FindAllSubjects {
	return NewFindAllSubjects(service.MakeFindAllSubjects())
}

func MakeDeleteSubject() DeleteSubject {
	return NewDeleteSubject(service.MakeDeleteSubject())
}
