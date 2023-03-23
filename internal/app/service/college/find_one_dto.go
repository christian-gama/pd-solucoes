package service

type FindOneCollegeInput struct {
	ID uint `validate:"required" uri:"id"`
}

type FindOneCollegeOutput struct {
	ID            uint                          `json:"id"`
	Name          string                        `json:"name"`
	Cnpj          string                        `json:"cnpj"`
	Courses       []*findOneCollegeCourseOutput `json:"courses"`
	TotalStudents int                           `json:"totalStudents"`
}

type findOneCollegeCourseOutput struct {
	ID   uint   `json:"id"   faker:"uint"`
	Name string `json:"name" faker:"len=50"`
}
