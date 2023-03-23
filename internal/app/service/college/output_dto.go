package service

type Output struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Cnpj         string          `json:"cnpj"`
	Courses      []*courseOutput `json:"courses"`
	StudentCount int             `json:"studentCount"`
}

type courseOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
