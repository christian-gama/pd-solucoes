package schema

// Teacher represents the teacher database schema.
type Teacher struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Degree   string
	Subjects []*Subject `gorm:"foreignKey:TeacherID"`
}

// TableName returns the table name.
func (Teacher) TableName() string {
	return "teachers"
}
