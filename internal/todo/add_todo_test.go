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
	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	testLogger = l.Sugar()
}

func TestAddTodoSuccess(t *testing.T) {
	creator := mocks.NewCreator(t)
	cat := time.Now()
	tid := uuid.NewString()
	expectedTodo := &todo.Todo{
		ID:          tid,
		Title:       "name",
		Description: "desc",
		CreatedAt:   cat,
		Completed:   false,
	}
	creator.On("CreateTodo", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), "name", "desc").Return(
		expectedTodo,
		nil,
	).Times(1)

	addTodo := todo.NewAddTodo(creator, testLogger)

	result, err := addTodo.Execute(context.TODO(), "name", "desc")

	assert.NoError(t, err)
	assert.EqualValues(t, expectedTodo, result)
	creator.AssertExpectations(t)
}

func TestAddTodoErrorMissingName(t *testing.T) {
	creator := mocks.NewCreator(t)
	addTodo := todo.NewAddTodo(creator, testLogger)

	result, err := addTodo.Execute(context.TODO(), "", "desc")
	assert.Nil(t, result)
	assert.EqualError(t, err, "could not create new todo, missing name")
}

func TestAddTodoErrorMissingDescription(t *testing.T) {
	creator := mocks.NewCreator(t)
	addTodo := todo.NewAddTodo(creator, testLogger)

	result, err := addTodo.Execute(context.TODO(), "name", "")
	assert.Nil(t, result)
	assert.EqualError(t, err, "could not create new todo, missing description")
}

func TestAddTodoErrorDuringCreation(t *testing.T) {
	creator := mocks.NewCreator(t)
	expectedErr := fmt.Errorf("the reason is you")
	creator.On("CreateTodo", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), "name", "desc").Return(
		nil,
		expectedErr,
	).Times(1)

	addTodo := todo.NewAddTodo(creator, testLogger)

	result, err := addTodo.Execute(context.TODO(), "name", "desc")
	assert.Nil(t, result)
	assert.ErrorIs(t, err, expectedErr)

	creator.AssertExpectations(t)
}
