package routes

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/internal/infra/router/routing"
	collegeCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/college"
	courseCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/course"
	courseSubjectCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/courseSubject"
	enrollmentCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/enrollment"
	studentCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/student"
	subjectCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/subject"
	teacherCtrl "github.com/christian-gama/pd-solucoes/internal/presentation/controller/teacher"
	"github.com/christian-gama/pd-solucoes/internal/presentation/middleware"
)

func Global() *routing.Routing {
	return &routing.Routing{
		Middlewares: []http.Middleware{
			middleware.MakeError(),
		},
	}
}

func Colleges() *routing.Routing {
	return &routing.Routing{
		Group: "/colleges",
		Routes: []*routing.Route{
			{Controller: collegeCtrl.MakeCreateCollege()},
			{Controller: collegeCtrl.MakeUpdateCollege()},
			{Controller: collegeCtrl.MakeFindOneCollege()},
			{Controller: collegeCtrl.MakeFindAllColleges()},
			{Controller: collegeCtrl.MakeDeleteCollege()},
		},
	}
}

func Teachers() *routing.Routing {
	return &routing.Routing{
		Group: "/teachers",
		Routes: []*routing.Route{
			{Controller: teacherCtrl.MakeCreateTeacher()},
			{Controller: teacherCtrl.MakeUpdateTeacher()},
			{Controller: teacherCtrl.MakeFindOneTeacher()},
			{Controller: teacherCtrl.MakeFindAllTeachers()},
			{Controller: teacherCtrl.MakeDeleteTeacher()},
		},
	}
}

func Students() *routing.Routing {
	return &routing.Routing{
		Group: "/students",
		Routes: []*routing.Route{
			{Controller: studentCtrl.MakeCreateStudent()},
			{Controller: studentCtrl.MakeUpdateStudent()},
			{Controller: studentCtrl.MakeFindOneStudent()},
			{Controller: studentCtrl.MakeFindAllStudents()},
			{Controller: studentCtrl.MakeDeleteStudent()},
		},
	}
}

func Courses() *routing.Routing {
	return &routing.Routing{
		Group: "/courses",
		Routes: []*routing.Route{
			{Controller: courseCtrl.MakeCreateCourse()},
			{Controller: courseCtrl.MakeUpdateCourse()},
			{Controller: courseCtrl.MakeFindOneCourse()},
			{Controller: courseCtrl.MakeFindAllCourses()},
			{Controller: courseCtrl.MakeDeleteCourse()},
		},
	}
}

func Subjects() *routing.Routing {
	return &routing.Routing{
		Group: "/subjects",
		Routes: []*routing.Route{
			{Controller: subjectCtrl.MakeCreateSubject()},
			{Controller: subjectCtrl.MakeUpdateSubject()},
			{Controller: subjectCtrl.MakeFindOneSubject()},
			{Controller: subjectCtrl.MakeFindAllSubjects()},
			{Controller: subjectCtrl.MakeDeleteSubject()},
		},
	}
}

func Enrollments() *routing.Routing {
	return &routing.Routing{
		Group: "/enrollments",
		Routes: []*routing.Route{
			{Controller: enrollmentCtrl.MakeCreateCourseEnrollment()},
			{Controller: enrollmentCtrl.MakeUpdateCourseEnrollment()},
			{Controller: enrollmentCtrl.MakeFindOneCourseEnrollment()},
			{Controller: enrollmentCtrl.MakeFindAllCourseEnrollments()},
			{Controller: enrollmentCtrl.MakeDeleteCourseEnrollment()},
		},
	}
}

func CourseSubjects() *routing.Routing {
	return &routing.Routing{
		Group: "/course-subjects",
		Routes: []*routing.Route{
			{Controller: courseSubjectCtrl.MakeCreateCourseSubject()},
			{Controller: courseSubjectCtrl.MakeUpdateCourseSubject()},
			{Controller: courseSubjectCtrl.MakeFindOneCourseSubject()},
			{Controller: courseSubjectCtrl.MakeFindAllCourseSubjects()},
			{Controller: courseSubjectCtrl.MakeDeleteCourseSubject()},
		},
	}
}
