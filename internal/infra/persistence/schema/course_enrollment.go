package schema

import "time"

// CourseEnrollment represents the courseEnrollment database schema.
type CourseEnrollment struct {
	ID              uint `gorm:"primaryKey"`
	StudentID       uint
	Student         *Student `gorm:"foreignKey:StudentID"`
	EnrollmentDate  time.Time
	CourseSubjectID uint
	CourseSubject   *CourseSubject `gorm:"foreignKey:CourseSubjectID"`
}

// TableName returns the table name.
func (CourseEnrollment) TableName() string {
	return "course_enrollments"
}
