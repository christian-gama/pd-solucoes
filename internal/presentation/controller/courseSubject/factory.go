package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"

func MakeCreateCourseSubject() CreateCourseSubject {
	return NewCreateCourseSubject(service.MakeCreateCourseSubject())
}

func MakeUpdateCourseSubject() UpdateCourseSubject {
	return NewUpdateCourseSubject(service.MakeUpdateCourseSubject())
}

func MakeFindOneCourseSubject() FindOneCourseSubject {
	return NewFindOneCourseSubject(service.MakeFindOneCourseSubject())
}

func MakeFindAllCourseSubjects() FindAllCourseSubjects {
	return NewFindAllCourseSubjects(service.MakeFindAllCourseSubjects())
}

func MakeDeleteCourseSubject() DeleteCourseSubject {
	return NewDeleteCourseSubject(service.MakeDeleteCourseSubject())
}
