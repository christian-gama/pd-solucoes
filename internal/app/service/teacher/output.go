package service

type Output struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Degree   string           `json:"degree"`
	Subjects []*subjectOutput `json:"subjects"`
}

type subjectOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
