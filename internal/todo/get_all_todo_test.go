package todo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/CDimonaco/todo-grpc-talk/internal/todo"
	"github.com/CDimonaco/todo-grpc-talk/internal/todo/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllTodosError(t *testing.T) {
	allTodoGetter := mocks.NewAllTodoGetter(t)
	allErr := fmt.Errorf("the reason is you")

	allTodoGetter.On(
		"GetAllTodos",
		mock.AnythingOfType("*context.emptyCtx"),
	).Return(nil, allErr).Times(1)

	getAllTodo := todo.NewGetAllTodo(testLogger, allTodoGetter)

	result, err := getAllTodo.Execute(context.TODO())
	assert.ErrorIs(t, err, allErr)
	assert.Nil(t, result)
	allTodoGetter.AssertExpectations(t)
}

func TestGetAllTodosSuccess(t *testing.T) {
	allTodoGetter := mocks.NewAllTodoGetter(t)
	expectedTodos := []todo.Todo{
		{
			ID:          uuid.NewString(),
			Title:       "name",
			Description: "desc",
			CreatedAt:   time.Now(),
			Completed:   false,
		},
		{
			ID:          uuid.NewString(),
			Title:       "name2",
			Description: "desc2",
			CreatedAt:   time.Now(),
			Completed:   false,
		},
	}
	allTodoGetter.On(
		"GetAllTodos",
		mock.AnythingOfType("*context.emptyCtx"),
	).Return(expectedTodos, nil).Times(1)

	getAllTodo := todo.NewGetAllTodo(testLogger, allTodoGetter)

	result, err := getAllTodo.Execute(context.TODO())
	assert.NoError(t, err)
	assert.EqualValues(t, expectedTodos, result)
	allTodoGetter.AssertExpectations(t)
}
