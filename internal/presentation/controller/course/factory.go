package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/course"

func MakeCreateCourse() CreateCourse {
	return NewCreateCourse(service.MakeCreateCourse())
}

func MakeUpdateCourse() UpdateCourse {
	return NewUpdateCourse(service.MakeUpdateCourse())
}

func MakeFindOneCourse() FindOneCourse {
	return NewFindOneCourse(service.MakeFindOneCourse())
}

func MakeFindAllCourses() FindAllCourses {
	return NewFindAllCourses(service.MakeFindAllCourses())
}

func MakeDeleteCourse() DeleteCourse {
	return NewDeleteCourse(service.MakeDeleteCourse())
}
