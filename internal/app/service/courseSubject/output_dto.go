package service

type Output struct {
	ID       uint             `json:"id"`
	Subject  *subjectOutput   `json:"subject"`
	Course   *courseOutput    `json:"course"`
	Students []*studentOutput `json:"students"`
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

type studentOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}
