package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// Course is a model that represents a course of a college.
type Course struct {
	ID        uint   `faker:"uint"`
	Name      string `faker:"len=50"`
	CollegeID uint   `faker:"uint"`
}

// NewCourse creates a new Course.
func NewCourse(id uint, name string, collegeID uint) (*Course, error) {
	m := &Course{
		ID:        id,
		Name:      name,
		CollegeID: collegeID,
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
