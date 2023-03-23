package schema

// Subject represents the subject database schema.
type Subject struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	TeacherID uint
	Teacher   *Teacher         `gorm:"foreignKey:TeacherID"`
	Courses   []*CourseSubject `gorm:"foreignKey:SubjectID"`
}

// TableName returns the table name.
func (Subject) TableName() string {
	return "subjects"
}
