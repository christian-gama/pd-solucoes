package service

import "time"

type Output struct {
	ID             uint                 `json:"id"`
	EnrollmentDate time.Time            `json:"enrollmentDate"`
	CourseSubject  *courseSubjectOutput `json:"courseSubject"`
}

type courseSubjectOutput struct {
	ID      uint           `json:"id"`
	Course  *courseOutput  `json:"course"`
	Subject *subjectOutput `json:"subject"`
}

type courseOutput struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CollegeID uint   `json:"collegeID"`
}

type subjectOutput struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	TeacherID uint   `json:"teacherID"`
}
