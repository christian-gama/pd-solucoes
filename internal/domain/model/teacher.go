package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

type Teacher struct {
	ID     uint   `faker:"uint"`
	Name   string `faker:"len=50"`
	Degree string `faker:"len=50"`
}

func NewTeacher(id uint, name string, degree string) (*Teacher, error) {
	m := &Teacher{
		ID:     id,
		Name:   name,
		Degree: degree,
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

	if errs.HasErrors() {
		return errs
	}

	return nil
}
