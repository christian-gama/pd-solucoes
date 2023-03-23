package schema

import (
	"gorm.io/gorm"
)

// College represents the college database schema.
type College struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Cnpj         string
	Courses      []*Course `gorm:"foreignKey:CollegeID"`
	StudentCount int       `gorm:"-"`
}

func (c *College) AfterFind(tx *gorm.DB) (err error) {
	var studentCount int64
	err = tx.Model(&CourseEnrollment{}).
		Joins("JOIN course_subjects ON course_subjects.id = course_enrollments.course_subject_id").
		Joins("JOIN courses ON courses.id = course_subjects.course_id").
		Where("courses.college_id = ?", c.ID).
		Distinct("course_enrollments.student_id").
		Count(&studentCount).Error

	if err != nil {
		return err
	}

	c.StudentCount = int(studentCount)
	return nil
}

// TableName returns the table name.
func (College) TableName() string {
	return "colleges"
}
