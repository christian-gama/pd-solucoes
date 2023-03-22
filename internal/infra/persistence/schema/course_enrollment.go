package schema

import "time"

// CourseEnrollment represents the courseEnrollment database schema.
type CourseEnrollment struct {
	ID              uint
	StudentID       uint
	Student         *Student
	CourseSubjectID uint
	CourseSubject   *CourseSubject
	EnrollmentDate  time.Time
}

// TableName returns the table name.
func (CourseEnrollment) TableName() string {
	return "course_enrollments"
}
