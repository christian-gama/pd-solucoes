package persistence

import (
	"github.com/christian-gama/pd-solucoes/internal/domain/repo"
	"github.com/christian-gama/pd-solucoes/internal/infra/sql"
)

func MakeCollege() repo.College {
	return NewCollege(sql.MakePostgres())
}

func MakeCourseEnrollment() repo.CourseEnrollment {
	return NewCourseEnrollment(sql.MakePostgres())
}

func MakeCourseSubject() repo.CourseSubject {
	return NewCourseSubject(sql.MakePostgres())
}

func MakeCourse() repo.Course {
	return NewCourse(sql.MakePostgres())
}

func MakeStudent() repo.Student {
	return NewStudent(sql.MakePostgres())
}

func MakeSubject() repo.Subject {
	return NewSubject(sql.MakePostgres())
}

func MakeTeacher() repo.Teacher {
	return NewTeacher(sql.MakePostgres())
}
