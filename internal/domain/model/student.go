package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// Student is the model that contains the student information.
type Student struct {
	ID   uint   `faker:"uint"`
	Name string `faker:"len=50"`
	Cpf  string `faker:"cpf"`
}

// NewStudent creates a new Student.
func NewStudent(id uint, name string, cpf string) (*Student, error) {
	m := &Student{
		ID:   id,
		Name: name,
		Cpf:  cpf,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

// Validate validates the Student.
func (m *Student) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.Cpf == "" {
		errs = errutil.Append(errs, errors.New("cpf is required"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
