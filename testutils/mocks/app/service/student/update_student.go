// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/student"
)

// UpdateStudent is an autogenerated mock type for the UpdateStudent type
type UpdateStudent struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *UpdateStudent) Handle(ctx context.Context, input *service.UpdateStudentInput) (*model.Student, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Student
	if rf, ok := ret.Get(0).(func(context.Context, *service.UpdateStudentInput) *model.Student); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.UpdateStudentInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUpdateStudent interface {
	mock.TestingT
	Cleanup(func())
}

// NewUpdateStudent creates a new instance of UpdateStudent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUpdateStudent(t mockConstructorTestingTNewUpdateStudent) *UpdateStudent {
	mock := &UpdateStudent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}