package service

type CreateCourseInput struct {
	Name      string `json:"name"      validate:"required,max=100,min=2" faker:"len=50"`
	CollegeID uint   `json:"collegeID" validate:"required,numeric"       faker:"uint"`
}
