// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/christian-gama/pd-solucoes/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"

	repo "github.com/christian-gama/pd-solucoes/internal/domain/repo"
)

// College is an autogenerated mock type for the College type
type College struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, params
func (_m *College) Create(ctx context.Context, params repo.CreateCollegeParams) (*model.College, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.College
	if rf, ok := ret.Get(0).(func(context.Context, repo.CreateCollegeParams) *model.College); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.College)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.CreateCollegeParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *College) Delete(ctx context.Context, params repo.DeleteCollegeParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.DeleteCollegeParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, params, preload
func (_m *College) FindAll(ctx context.Context, params repo.FindAllCollegeParams, preload ...string) (*querying.PaginationOutput[*model.College], error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *querying.PaginationOutput[*model.College]
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindAllCollegeParams, ...string) *querying.PaginationOutput[*model.College]); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*model.College])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindAllCollegeParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, params, preload
func (_m *College) FindOne(ctx context.Context, params repo.FindOneCollegeParams, preload ...string) (*model.College, error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.College
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindOneCollegeParams, ...string) *model.College); ok {
		r0 = rf(ctx, params, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.College)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.FindOneCollegeParams, ...string) error); ok {
		r1 = rf(ctx, params, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *College) Update(ctx context.Context, params repo.UpdateCollegeParams) (*model.College, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.College
	if rf, ok := ret.Get(0).(func(context.Context, repo.UpdateCollegeParams) *model.College); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.College)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repo.UpdateCollegeParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCollege interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollege creates a new instance of College. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollege(t mockConstructorTestingTNewCollege) *College {
	mock := &College{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
