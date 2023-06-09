// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	mock "github.com/stretchr/testify/mock"
)

// Filterer is an autogenerated mock type for the Filterer type
type Filterer struct {
	mock.Mock
}

// Add provides a mock function with given fields: field, operator, value
func (_m *Filterer) Add(field string, operator string, value interface{}) querying.Filterer {
	ret := _m.Called(field, operator, value)

	var r0 querying.Filterer
	if rf, ok := ret.Get(0).(func(string, string, interface{}) querying.Filterer); ok {
		r0 = rf(field, operator, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(querying.Filterer)
		}
	}

	return r0
}

// Field provides a mock function with given fields: idx
func (_m *Filterer) Field(idx int) string {
	ret := _m.Called(idx)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Operator provides a mock function with given fields: idx
func (_m *Filterer) Operator(idx int) string {
	ret := _m.Called(idx)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Slice provides a mock function with given fields:
func (_m *Filterer) Slice() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Value provides a mock function with given fields: idx
func (_m *Filterer) Value(idx int) string {
	ret := _m.Called(idx)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewFilterer interface {
	mock.TestingT
	Cleanup(func())
}

// NewFilterer creates a new instance of Filterer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFilterer(t mockConstructorTestingTNewFilterer) *Filterer {
	mock := &Filterer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
