package service

type Output struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Cnpj         string          `json:"cnpj"`
	Courses      []*courseOutput `json:"courses"`
	StudentCount int             `json:"studentCount"`
}

type courseOutput struct {
	ID   uint   `json:"id"   faker:"uint"`
	Name string `json:"name" faker:"len=50"`
}

type PaginationOutput struct {
	Total   int       `json:"total"`
	Results []*Output `json:"results"`
}
