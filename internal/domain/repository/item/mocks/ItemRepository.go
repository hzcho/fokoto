// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	item "fokoto/internal/domain/model/item"

	mock "github.com/stretchr/testify/mock"
)

// ItemRepository is an autogenerated mock type for the ItemRepository type
type ItemRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: orderId
func (_m *ItemRepository) Get(orderId uint64) ([]item.Item, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []item.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]item.Item, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(uint64) []item.Item); ok {
		r0 = rf(orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]item.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAll provides a mock function with given fields: orderId, items
func (_m *ItemRepository) SaveAll(orderId uint64, items []item.Item) error {
	ret := _m.Called(orderId, items)

	if len(ret) == 0 {
		panic("no return value specified for SaveAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, []item.Item) error); ok {
		r0 = rf(orderId, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewItemRepository creates a new instance of ItemRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewItemRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ItemRepository {
	mock := &ItemRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
