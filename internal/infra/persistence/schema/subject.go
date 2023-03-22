package schema

// Subject represents the subject database schema.
type Subject struct {
	ID        uint
	Name      string
	TeacherID uint
	Teacher   *Teacher
}

// TableName returns the table name.
func (Subject) TableName() string {
	return "subjects"
}
