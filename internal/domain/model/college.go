package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

type College struct {
	ID   uint   `faker:"uint"`
	Name string `faker:"len=50"`
	Cnpj string `faker:"cnpj"`
}

func NewCollege(id uint, name, cnpj string) (*College, error) {
	m := &College{
		ID:   id,
		Name: name,
		Cnpj: cnpj,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

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
