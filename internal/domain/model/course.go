package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// Course is a model that represents a course of a college.
type Course struct {
	ID          uint                `json:"id,omitempty"          faker:"uint"`
	Name        string              `json:"name,omitempty"        faker:"len=50"`
	CollegeID   uint                `json:"collegeID,omitempty"   faker:"uint"`
	College     *College            `json:"college,omitempty"     faker:"-"`
	Subjects    []*CourseSubject    `json:"subject,omitempty"     faker:"-"`
	Enrollments []*CourseEnrollment `json:"enrollments,omitempty" faker:"-"`
}

// NewCourse creates a new Course.
func NewCourse(id uint, name string, collegeID uint) (*Course, error) {
	m := &Course{
		ID:          id,
		Name:        name,
		CollegeID:   collegeID,
		College:     nil,
		Subjects:    nil,
		Enrollments: nil,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

// Validate validates the Course.
func (m *Course) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.CollegeID == 0 {
		errs = errutil.Append(errs, errors.New("collegeID is required"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
