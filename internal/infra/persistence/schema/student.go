package schema

// Student represents the student database schema.
type Student struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Cpf  string
}

// TableName returns the table name.
func (Student) TableName() string {
	return "students"
}
