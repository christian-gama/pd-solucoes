package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// Subject is the model of a subject.
type Subject struct {
	ID        uint             `json:"id,omitempty"        faker:"uint"`
	Name      string           `json:"name,omitempty"      faker:"len=50"`
	TeacherID uint             `json:"teacherID,omitempty" faker:"uint"`
	Teacher   *Teacher         `json:"teacher,omitempty"   faker:"-"`
	Courses   []*CourseSubject `json:"courses,omitempty"   faker:"-"`
}

func NewSubject(id uint, name string, teacherID uint) (*Subject, error) {
	m := &Subject{
		ID:        id,
		Name:      name,
		TeacherID: teacherID,
		Teacher:   nil,
		Courses:   nil,
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

	if errs.HasErrors() {
		return errs
	}

	return nil
}
