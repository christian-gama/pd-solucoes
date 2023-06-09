// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"

	repo "github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

// Student is an autogenerated mock type for the Student type
type Student struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, params
func (_m *Student) Create(ctx context.Context, params repo.CreateStudentParams) (*model.Student, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.Student
	if rf, ok := ret.Get(0).(func(context.Context, repo.CreateStudentParams) *model.Student); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.CreateStudentParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *Student) Delete(ctx context.Context, params repo.DeleteStudentParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.DeleteStudentParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, params, preload
func (_m *Student) FindAll(ctx context.Context, params repo.FindAllStudentParams, preload ...string) (*querying.PaginationOutput[*model.Student], error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *querying.PaginationOutput[*model.Student]
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindAllStudentParams, ...string) *querying.PaginationOutput[*model.Student]); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*model.Student])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindAllStudentParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, params, preload
func (_m *Student) FindOne(ctx context.Context, params repo.FindOneStudentParams, preload ...string) (*model.Student, error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Student
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindOneStudentParams, ...string) *model.Student); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindOneStudentParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *Student) Update(ctx context.Context, params repo.UpdateStudentParams) (*model.Student, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.Student
	if rf, ok := ret.Get(0).(func(context.Context, repo.UpdateStudentParams) *model.Student); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.UpdateStudentParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStudent interface {
	mock.TestingT
	Cleanup(func())
}

// NewStudent creates a new instance of Student. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStudent(t mockConstructorTestingTNewStudent) *Student {
	mock := &Student{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
