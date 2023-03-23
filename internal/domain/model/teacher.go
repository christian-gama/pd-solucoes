package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// Teacher is the model that contains the teacher information.
type Teacher struct {
	ID       uint       `json:"id,omitempty"       faker:"uint"`
	Name     string     `json:"name,omitempty"     faker:"len=50"`
	Degree   string     `json:"degree,omitempty"   faker:"len=50"`
	Subjects []*Subject `json:"subjects,omitempty" faker:"-"`
}

// NewTeacher creates a new Teacher.
func NewTeacher(id uint, name string, degree string) (*Teacher, error) {
	m := &Teacher{
		ID:       id,
		Name:     name,
		Degree:   degree,
		Subjects: nil,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

// Validate validates the Teacher.
func (m *Teacher) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.Degree == "" {
		errs = errutil.Append(errs, errors.New("degree is required"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
