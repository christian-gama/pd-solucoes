// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	mock "github.com/stretchr/testify/mock"
)

// CreateSubject is an autogenerated mock type for the CreateSubject type
type CreateSubject struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *CreateSubject) Handle(ctx context.Context, input *service.CreateSubjectInput) (*service.Output, error) {
	ret := _m.Called(ctx, input)

	var r0 *service.Output
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateSubjectInput) *service.Output); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateSubjectInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCreateSubject interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateSubject creates a new instance of CreateSubject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateSubject(t mockConstructorTestingTNewCreateSubject) *CreateSubject {
	mock := &CreateSubject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}