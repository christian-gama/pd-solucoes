// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	mock "github.com/stretchr/testify/mock"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/college"
)

// FindAllColleges is an autogenerated mock type for the FindAllColleges type
type FindAllColleges struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *FindAllColleges) Handle(ctx context.Context, input *service.FindAllInput) (*querying.PaginationOutput[*service.Output], error) {
	ret := _m.Called(ctx, input)

	var r0 *querying.PaginationOutput[*service.Output]
	if rf, ok := ret.Get(0).(func(context.Context, *service.FindAllInput) *querying.PaginationOutput[*service.Output]); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*service.Output])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FindAllInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindAllColleges interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAllColleges creates a new instance of FindAllColleges. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAllColleges(t mockConstructorTestingTNewFindAllColleges) *FindAllColleges {
	mock := &FindAllColleges{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
