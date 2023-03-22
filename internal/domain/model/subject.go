package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

type Subject struct {
	ID        uint   `faker:"uint"`
	Name      string `faker:"len=50"`
	TeacherID uint   `faker:"uint"`
	Teacher   *Teacher
}

func NewSubject(id uint, name string, teacherID uint, teacher *Teacher) (*Subject, error) {
	m := &Subject{
		ID:        id,
		Name:      name,
		TeacherID: teacherID,
		Teacher:   teacher,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Subject) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.TeacherID == 0 {
		errs = errutil.Append(errs, errors.New("teacher is required"))
	}

	if m.Teacher == nil {
		errs = errutil.Append(errs, errors.New("teacher is required"))
	} else if err := m.Teacher.Validate(); err != nil {
		errs = errutil.Append(errs, errors.New("teacher is invalid"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
