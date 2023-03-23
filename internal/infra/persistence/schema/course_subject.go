package schema

// CourseSubject is the join table between courses and subjects.
type CourseSubject struct {
	ID        uint `gorm:"primaryKey"`
	CourseID  uint
	SubjectID uint
	Course    *Course    `gorm:"foreignKey:CourseID"`
	Subject   *Subject   `gorm:"foreignKey:SubjectID"`
	Students  []*Student `gorm:"many2many:course_enrollments;"`
}

// TableName returns the table name.
func (CourseSubject) TableName() string {
	return "course_subjects"
}
