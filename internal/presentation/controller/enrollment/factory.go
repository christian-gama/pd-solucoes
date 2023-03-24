package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"

func MakeCreateCourseEnrollment() CreateCourseEnrollment {
	return NewCreateCourseEnrollment(service.MakeCreateCourseEnrollment())
}

func MakeUpdateCourseEnrollment() UpdateCourseEnrollment {
	return NewUpdateCourseEnrollment(service.MakeUpdateCourseEnrollment())
}

func MakeFindOneCourseEnrollment() FindOneCourseEnrollment {
	return NewFindOneCourseEnrollment(service.MakeFindOneCourseEnrollment())
}

func MakeFindAllCourseEnrollments() FindAllCourseEnrollments {
	return NewFindAllCourseEnrollments(service.MakeFindAllCourseEnrollments())
}

func MakeDeleteCourseEnrollment() DeleteCourseEnrollment {
	return NewDeleteCourseEnrollment(service.MakeDeleteCourseEnrollment())
}
