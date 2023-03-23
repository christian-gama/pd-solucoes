package service

type Output struct {
	ID             uint                   `json:"id"`
	Name           string                 `json:"name"`
	Cpf            string                 `json:"cpf"`
	CourseSubjects []*courseSubjectOutput `json:"courseSubjects"`
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
