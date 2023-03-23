package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// College is the model of a college.
type College struct {
	ID           uint      `faker:"uint"`
	Name         string    `faker:"len=50"`
	Cnpj         string    `faker:"cnpj"`
	Courses      []*Course `faker:"-"`
	StudentCount int       `faker:"-"`
}

// NewCollege creates a new College.
func NewCollege(id uint, name, cnpj string) (*College, error) {
	m := &College{
		ID:           id,
		Name:         name,
		Cnpj:         cnpj,
		Courses:      make([]*Course, 0),
		StudentCount: 0,
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

	if errs.HasErrors() {
		return errs
	}

	return nil
}
