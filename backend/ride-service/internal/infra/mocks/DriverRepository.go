// Code generated by mockery v2.31.1. DO NOT EDIT.

package mocks

import (
	"github.com/chjoaquim/ride-service/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// DriverRepository is an autogenerated mock type for the DriverRepository type
type DriverRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: driver
func (_m *DriverRepository) Create(driver *domain.Driver) (*domain.Driver, error) {
	ret := _m.Called(driver)

	var r0 *domain.Driver
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Driver) (*domain.Driver, error)); ok {
		return rf(driver)
	}
	if rf, ok := ret.Get(0).(func(*domain.Driver) *domain.Driver); ok {
		r0 = rf(driver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Driver)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Driver) error); ok {
		r1 = rf(driver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *DriverRepository) Get(id string) (*domain.Driver, error) {
	ret := _m.Called(id)

	var r0 *domain.Driver
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Driver, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Driver); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Driver)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDriverRepository creates a new instance of DriverRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDriverRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DriverRepository {
	mock := &DriverRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}