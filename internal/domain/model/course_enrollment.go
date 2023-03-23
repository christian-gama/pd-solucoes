package model

import (
	"errors"
	"time"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// CourseEnrollment is the model of a course enrollment.
type CourseEnrollment struct {
	ID              uint `faker:"uint"`
	StudentID       uint `faker:"uint"`
	Student         *Student
	EnrollmentDate  time.Time `faker:"time_now"`
	CourseSubjectID uint      `faker:"uint"`
	CourseSubject   *CourseSubject
}

func NewCourseEnrollment(
	id uint,
	studentID uint,
	student *Student,
	enrollmentDate time.Time,
	courseSubjectID uint,
	courseSubject *CourseSubject,
) (*CourseEnrollment, error) {
	m := &CourseEnrollment{
		ID:              id,
		StudentID:       studentID,
		Student:         student,
		EnrollmentDate:  enrollmentDate,
		CourseSubjectID: courseSubjectID,
		CourseSubject:   courseSubject,
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

	if m.Student == nil {
		errs = errutil.Append(errs, errors.New("student is required"))
	} else if err := m.Student.Validate(); err != nil {
		errs = errutil.Append(errs, errors.New("student is invalid"))
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

	if m.CourseSubject == nil {
		errs = errutil.Append(errs, errors.New("course is required"))
	} else if err := m.CourseSubject.Validate(); err != nil {
		errs = errutil.Append(errs, errors.New("course is invalid"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
