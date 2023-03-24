// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	mock "github.com/stretchr/testify/mock"
)

// CreateCourseSubject is an autogenerated mock type for the CreateCourseSubject type
type CreateCourseSubject struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *CreateCourseSubject) Handle(ctx context.Context, input *service.CreateCourseSubjectInput) (*service.Output, error) {
	ret := _m.Called(ctx, input)

	var r0 *service.Output
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateCourseSubjectInput) *service.Output); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateCourseSubjectInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCreateCourseSubject interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateCourseSubject creates a new instance of CreateCourseSubject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateCourseSubject(t mockConstructorTestingTNewCreateCourseSubject) *CreateCourseSubject {
	mock := &CreateCourseSubject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}