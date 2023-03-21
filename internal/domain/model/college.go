package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// College is the model of a college.
type College struct {
	ID      uint   `faker:"uint"`
	Name    string `faker:"len=50"`
	Cnpj    string `faker:"cnpj"`
	Courses []*Course
}

// NewCollege creates a new College.
func NewCollege(id uint, name, cnpj string, courses []*Course) (*College, error) {
	m := &College{
		ID:      id,
		Name:    name,
		Cnpj:    cnpj,
		Courses: courses,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

// Validate validates the College.
func (m *College) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.Cnpj == "" {
		errs = errutil.Append(errs, errors.New("cnpj is required"))
	}

	if len(m.Courses) == 0 {
		errs = errutil.Append(errs, errors.New("courses is required"))
	} else {
		for _, course := range m.Courses {
			if err := course.Validate(); err != nil {
				errs = errutil.Append(errs, err)
			}
		}
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
