package service

import "time"

type Output struct {
	ID             uint
	EnrollmentDate time.Time
	CourseSubject  *courseSubjectOutput
}

type courseSubjectOutput struct {
	ID      uint           `json:"id"`
	Course  *courseOutput  `json:"course"`
	Subject *subjectOutput `json:"subject"`
}

type studentOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
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
