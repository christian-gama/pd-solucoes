package schema

// College represents the college database schema.
type College struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Cnpj    string
	Courses []*Course `gorm:"foreignKey:CollegeID"`
}

// TableName returns the table name.
func (College) TableName() string {
	return "colleges"
}
