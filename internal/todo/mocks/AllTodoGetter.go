// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	todo "github.com/CDimonaco/todo-grpc-talk/internal/todo"
	mock "github.com/stretchr/testify/mock"
)

// AllTodoGetter is an autogenerated mock type for the AllTodoGetter type
type AllTodoGetter struct {
	mock.Mock
}

// GetAllTodos provides a mock function with given fields: ctx
func (_m *AllTodoGetter) GetAllTodos(ctx context.Context) ([]todo.Todo, error) {
	ret := _m.Called(ctx)

	var r0 []todo.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []todo.Todo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]todo.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAllTodoGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAllTodoGetter creates a new instance of AllTodoGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAllTodoGetter(t mockConstructorTestingTNewAllTodoGetter) *AllTodoGetter {
	mock := &AllTodoGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}