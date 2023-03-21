package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

type Teacher struct {
	ID       uint   `faker:"uint"`
	Name     string `faker:"len=50"`
	Degree   string `faker:"len=50"`
	Subjects []*Subject
}

func NewTeacher(id uint, name string, degree string, subjects []*Subject) (*Teacher, error) {
	m := &Teacher{
		ID:       id,
		Name:     name,
		Degree:   degree,
		Subjects: subjects,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Teacher) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.Degree == "" {
		errs = errutil.Append(errs, errors.New("degree is required"))
	}

	if m.Subjects == nil {
		errs = errutil.Append(errs, errors.New("subjects is required"))
	} else {
		for _, subject := range m.Subjects {
			if err := subject.Validate(); err != nil {
				errs = errutil.Append(errs, err)
			}
		}
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
