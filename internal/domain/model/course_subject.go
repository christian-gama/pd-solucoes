package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// CourseSubject is the model that contains the relationship between courses and subjects.
type CourseSubject struct {
	ID        uint       `json:"id,omitempty"        faker:"uint"`
	CourseID  uint       `json:"courseID,omitempty"  faker:"uint"`
	SubjectID uint       `json:"subjectID,omitempty" faker:"uint"`
	Course    *Course    `json:"course,omitempty"    faker:"-"`
	Subject   *Subject   `json:"subject,omitempty"   faker:"-"`
	Students  []*Student `json:"students,omitempty"  faker:"-"`
}

// NewCourseSubject creates a new CourseSubject.
func NewCourseSubject(id, courseID, subjectID uint,
) (*CourseSubject, error) {
	m := &CourseSubject{
		ID:        id,
		CourseID:  courseID,
		SubjectID: subjectID,
		Course:    nil,
		Subject:   nil,
		Students:  nil,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

// Validate validates the CourseSubject.
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
