// Code generated by mockery v2.9.6. DO NOT EDIT.

package util

import (
	mock "github.com/stretchr/testify/mock"
	netlink "github.com/vishvananda/netlink"
)

// MockLink is an autogenerated mock type for the Link type
type MockLink struct {
	mock.Mock
}

// Attrs provides a mock function with given fields:
func (_m *MockLink) Attrs() *netlink.LinkAttrs {
	ret := _m.Called()

	var r0 *netlink.LinkAttrs
	if rf, ok := ret.Get(0).(func() *netlink.LinkAttrs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*netlink.LinkAttrs)
		}
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *MockLink) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockLink interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockLink creates a new instance of MockLink. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockLink(t mockConstructorTestingTNewMockLink) *MockLink {
	mock := &MockLink{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
