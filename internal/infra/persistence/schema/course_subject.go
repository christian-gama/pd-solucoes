package schema

// CourseSubject is the join table between courses and subjects.
type CourseSubject struct {
	ID        uint `gorm:"primaryKey"`
	CourseID  uint
	SubjectID uint
	Course    *Course
	Subject   *Subject
}

// TableName returns the table name.
func (CourseSubject) TableName() string {
	return "course_subjects"
}
