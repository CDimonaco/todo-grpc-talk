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

func TestGetTodoErrorNoIdProvided(t *testing.T) {
	todoGetter := mocks.NewTodoGetter(t)

	getTodo := todo.NewGetTodo(testLogger, todoGetter)
	result, err := getTodo.Execute(context.TODO(), "")

	assert.EqualError(t, err, "empty id provided")
	assert.Nil(t, result)
}
func TestGetTodoErrorNoUUID(t *testing.T) {
	todoGetter := mocks.NewTodoGetter(t)

	getTodo := todo.NewGetTodo(testLogger, todoGetter)
	result, err := getTodo.Execute(context.TODO(), "theid")

	assert.ErrorContains(t, err, "the provided id theid is not an uuid")
	assert.Nil(t, result)
}

func TestGetTodoErrorGetter(t *testing.T) {
	todoGetter := mocks.NewTodoGetter(t)
	todoID := uuid.NewString()
	getErr := fmt.Errorf("the reason is you")
	todoGetter.On(
		"GetTodoByID",
		mock.AnythingOfType("*context.emptyCtx"),
		todoID,
	).Return(nil, getErr).Times(1)
	getTodo := todo.NewGetTodo(testLogger, todoGetter)
	result, err := getTodo.Execute(context.TODO(), todoID)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, getErr)

}

func TestGetTodoSuccess(t *testing.T) {
	todoGetter := mocks.NewTodoGetter(t)
	expectedTodo := &todo.Todo{
		ID:          uuid.NewString(),
		Title:       "name",
		Description: "desc",
		CreatedAt:   time.Now(),
	}

	todoGetter.On(
		"GetTodoByID",
		mock.AnythingOfType("*context.emptyCtx"),
		expectedTodo.ID,
	).Return(expectedTodo, nil).Times(1)

	getTodo := todo.NewGetTodo(testLogger, todoGetter)
	result, err := getTodo.Execute(context.TODO(), expectedTodo.ID)

	assert.NoError(t, err)
	assert.EqualValues(t, expectedTodo, result)
}
