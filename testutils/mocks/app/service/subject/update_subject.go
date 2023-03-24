// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/subject"
	mock "github.com/stretchr/testify/mock"
)

// UpdateSubject is an autogenerated mock type for the UpdateSubject type
type UpdateSubject struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *UpdateSubject) Handle(ctx context.Context, input *service.UpdateInput) (*service.CreateOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *service.CreateOutput
	if rf, ok := ret.Get(0).(func(context.Context, *service.UpdateInput) *service.CreateOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.UpdateInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUpdateSubject interface {
	mock.TestingT
	Cleanup(func())
}

// NewUpdateSubject creates a new instance of UpdateSubject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUpdateSubject(t mockConstructorTestingTNewUpdateSubject) *UpdateSubject {
	mock := &UpdateSubject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
