// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	todo "github.com/CDimonaco/todo-grpc-talk/internal/todo"
	mock "github.com/stretchr/testify/mock"
)

// TodoGetter is an autogenerated mock type for the TodoGetter type
type TodoGetter struct {
	mock.Mock
}

// GetTodoByID provides a mock function with given fields: ctx, id
func (_m *TodoGetter) GetTodoByID(ctx context.Context, id string) (*todo.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 *todo.Todo
	if rf, ok := ret.Get(0).(func(context.Context, string) *todo.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoGetter creates a new instance of TodoGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoGetter(t mockConstructorTestingTNewTodoGetter) *TodoGetter {
	mock := &TodoGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}