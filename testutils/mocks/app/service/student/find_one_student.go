// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
	mock "github.com/stretchr/testify/mock"
)

// FindOneStudent is an autogenerated mock type for the FindOneStudent type
type FindOneStudent struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *FindOneStudent) Handle(ctx context.Context, input *service.FindOneInput) (*service.Output, error) {
	ret := _m.Called(ctx, input)

	var r0 *service.Output
	if rf, ok := ret.Get(0).(func(context.Context, *service.FindOneInput) *service.Output); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FindOneInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindOneStudent interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindOneStudent creates a new instance of FindOneStudent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindOneStudent(t mockConstructorTestingTNewFindOneStudent) *FindOneStudent {
	mock := &FindOneStudent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
