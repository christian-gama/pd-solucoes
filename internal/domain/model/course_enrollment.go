package model

import (
	"errors"
	"time"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// CourseEnrollment is the model of a course enrollment.
type CourseEnrollment struct {
	ID              uint           `faker:"uint"`
	StudentID       uint           `faker:"uint"`
	Student         *Student       `faker:"-"`
	EnrollmentDate  time.Time      `faker:"time_now"`
	CourseSubjectID uint           `faker:"uint"`
	CourseSubject   *CourseSubject `faker:"-"`
}

func NewCourseEnrollment(
	id uint,
	studentID uint,
	enrollmentDate time.Time,
	courseSubjectID uint,
) (*CourseEnrollment, error) {
	m := &CourseEnrollment{
		ID:              id,
		StudentID:       studentID,
		Student:         nil,
		EnrollmentDate:  enrollmentDate,
		CourseSubjectID: courseSubjectID,
		CourseSubject:   nil,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *CourseEnrollment) Validate() error {
	var errs *errutil.Error

	if m.StudentID == 0 {
		errs = errutil.Append(errs, errors.New("student is required"))
	}

	if m.EnrollmentDate.IsZero() {
		errs = errutil.Append(errs, errors.New("enrollment date is required"))
	}

	if m.EnrollmentDate.After(time.Now()) {
		errs = errutil.Append(
			errs,
			errors.New("enrollment date is invalid, it cannot be in the future"),
		)
	}

	if m.CourseSubjectID == 0 {
		errs = errutil.Append(errs, errors.New("course is required"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
