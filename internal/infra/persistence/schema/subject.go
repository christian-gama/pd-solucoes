package schema

// Subject represents the subject database schema.
type Subject struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	TeacherID uint
}

// TableName returns the table name.
func (Subject) TableName() string {
	return "subjects"
}
