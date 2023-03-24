// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	querying "github.com/christian-gama/pd-solucoes/internal/domain/querying"
	mock "github.com/stretchr/testify/mock"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
)

// FindAllCourseEnrollments is an autogenerated mock type for the FindAllCourseEnrollments type
type FindAllCourseEnrollments struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *FindAllCourseEnrollments) Handle(ctx context.Context, input *service.FindAllCourseEnrollmentsInput) (*querying.PaginationOutput[*service.Output], error) {
	ret := _m.Called(ctx, input)

	var r0 *querying.PaginationOutput[*service.Output]
	if rf, ok := ret.Get(0).(func(context.Context, *service.FindAllCourseEnrollmentsInput) *querying.PaginationOutput[*service.Output]); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*service.Output])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.FindAllCourseEnrollmentsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFindAllCourseEnrollments interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAllCourseEnrollments creates a new instance of FindAllCourseEnrollments. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAllCourseEnrollments(t mockConstructorTestingTNewFindAllCourseEnrollments) *FindAllCourseEnrollments {
	mock := &FindAllCourseEnrollments{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}