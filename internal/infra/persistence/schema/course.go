package schema

// Course represents the course database schema.
type Course struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	CollegeID   uint
	College     *College            `gorm:"foreignKey:CollegeID"`
	Subjects    []*CourseSubject    `gorm:"foreignKey:CourseID"`
	Enrollments []*CourseEnrollment `gorm:"foreignKey:CourseSubjectID"`
}

// TableName returns the table name.
func (Course) TableName() string {
	return "courses"
}
