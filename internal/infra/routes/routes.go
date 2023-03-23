package routes

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/http"
	"github.com/christian-gama/pd-solucoes/internal/infra/router/routing"
	collegeController "github.com/christian-gama/pd-solucoes/internal/presentation/controller/college"
	courseController "github.com/christian-gama/pd-solucoes/internal/presentation/controller/course"
	studentController "github.com/christian-gama/pd-solucoes/internal/presentation/controller/student"
	subjectController "github.com/christian-gama/pd-solucoes/internal/presentation/controller/subject"
	teacherController "github.com/christian-gama/pd-solucoes/internal/presentation/controller/teacher"
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
			{Controller: collegeController.MakeCreateCollege()},
			{Controller: collegeController.MakeUpdateCollege()},
			{Controller: collegeController.MakeFindOneCollege()},
			{Controller: collegeController.MakeFindAllColleges()},
			{Controller: collegeController.MakeDeleteCollege()},
		},
	}
}

func Teachers() *routing.Routing {
	return &routing.Routing{
		Group: "/teachers",
		Routes: []*routing.Route{
			{Controller: teacherController.MakeCreateTeacher()},
			{Controller: teacherController.MakeUpdateTeacher()},
			{Controller: teacherController.MakeFindOneTeacher()},
			{Controller: teacherController.MakeFindAllTeachers()},
			{Controller: teacherController.MakeDeleteTeacher()},
		},
	}
}

func Students() *routing.Routing {
	return &routing.Routing{
		Group: "/students",
		Routes: []*routing.Route{
			{Controller: studentController.MakeCreateStudent()},
			{Controller: studentController.MakeUpdateStudent()},
			{Controller: studentController.MakeFindOneStudent()},
			{Controller: studentController.MakeFindAllStudents()},
			{Controller: studentController.MakeDeleteStudent()},
		},
	}
}

func Courses() *routing.Routing {
	return &routing.Routing{
		Group: "/courses",
		Routes: []*routing.Route{
			{Controller: courseController.MakeCreateCourse()},
			{Controller: courseController.MakeUpdateCourse()},
			{Controller: courseController.MakeFindOneCourse()},
			{Controller: courseController.MakeFindAllCourses()},
			{Controller: courseController.MakeDeleteCourse()},
		},
	}
}

func Subjects() *routing.Routing {
	return &routing.Routing{
		Group: "/subjects",
		Routes: []*routing.Route{
			{Controller: subjectController.MakeCreateSubject()},
			{Controller: subjectController.MakeUpdateSubject()},
			{Controller: subjectController.MakeFindOneSubject()},
			{Controller: subjectController.MakeFindAllSubjects()},
			{Controller: subjectController.MakeDeleteSubject()},
		},
	}
}
