package service

type CreateInput struct {
	Name      string `json:"name"      validate:"required,max=100,min=2" faker:"len=50"`
	CollegeID uint   `json:"collegeID" validate:"required,numeric"       faker:"uint"`
}

type CreateOutput struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CollegeID uint   `json:"collegeID"`
}
