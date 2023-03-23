package service

type Output struct {
	ID      uint                   `json:"id"`
	Name    string                 `json:"name"`
	Teacher *teacherOutput         `json:"teacher"`
	Courses []*courseSubjectOutput `json:"courses"`
}

type teacherOutput struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Degree string `json:"degree"`
}

type courseSubjectOutput struct {
	CourseID uint             `json:"courseID"`
	Students []*studentOutput `json:"students"`
}

type studentOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}
