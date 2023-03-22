package schema

// Course represents the course database schema.
type Course struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CollegeID uint
	College   *College
}

// TableName returns the table name.
func (Course) TableName() string {
	return "courses"
}
