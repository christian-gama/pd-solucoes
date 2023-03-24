package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateCourseEnrollment() CreateCourseEnrollment {
	return NewCreateCourseEnrollment(persistence.MakeCourseEnrollment())
}

func MakeUpdateCourseEnrollment() UpdateCourseEnrollment {
	return NewUpdateCourseEnrollment(persistence.MakeCourseEnrollment())
}

func MakeFindOneCourseEnrollment() FindOneCourseEnrollment {
	return NewFindOneCourseEnrollment(persistence.MakeCourseEnrollment())
}

func MakeFindAllCourseEnrollments() FindAllCourseEnrollments {
	return NewFindAllCourseEnrollments(persistence.MakeCourseEnrollment())
}

func MakeDeleteCourseEnrollment() DeleteCourseEnrollment {
	return NewDeleteCourseEnrollment(persistence.MakeCourseEnrollment())
}
