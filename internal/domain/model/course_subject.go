package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

type CourseSubject struct {
	ID        uint `faker:"uint"`
	CourseID  uint `faker:"uint"`
	SubjectID uint `faker:"uint"`
}

func NewCourseSubject(id, courseID, subjectID uint) (*CourseSubject, error) {
	m := &CourseSubject{
		ID:        id,
		CourseID:  courseID,
		SubjectID: subjectID,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *CourseSubject) Validate() error {
	var errs *errutil.Error

	if m.CourseID == 0 {
		errs = errutil.Append(errs, errors.New("courseID is required"))
	}

	if m.SubjectID == 0 {
		errs = errutil.Append(errs, errors.New("subjectID is required"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
