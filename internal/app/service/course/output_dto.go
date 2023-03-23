package service

type Output struct {
	ID        uint             `json:"id"`
	Name      string           `json:"name"`
	CollegeID uint             `json:"collegeID"`
	Subjects  []*subjectOutput `json:"subjects"`
}

type subjectOutput struct {
	ID        uint             `json:"id"`
	CourseID  uint             `json:"courseID"`
	SubjectID uint             `json:"subjectID"`
	Students  []*studentOutput `json:"students"`
}

type studentOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}
