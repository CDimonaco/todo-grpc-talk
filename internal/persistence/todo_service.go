package persistence

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/CDimonaco/todo-grpc-talk/internal/todo"
	"go.uber.org/zap"
)

type TodoService struct {
	m      sync.Map
	logger *zap.SugaredLogger
}

func NewTodoService(
	logger *zap.SugaredLogger,
) *TodoService {
	l := logger.With("component", "todoService")

	return &TodoService{
		m:      sync.Map{},
		logger: l,
	}

}

func (s *TodoService) CreateTodo(ctx context.Context, id string, name string, description string) (*todo.Todo, error) {
	t := todo.Todo{
		ID:          id,
		Title:       name,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}

	s.m.Store(id, t)

	return &t, nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, id string) (*todo.Todo, error) {
	item, ok := s.m.Load(id)
	if !ok {
		s.logger.Debugw("todo not found", "todo_id", id)
		return nil, nil
	}

	t, ok := item.(todo.Todo)
	if !ok {
		s.logger.Errorw("could not cast todo item from map", "todo_id", id, "raw_value", item)
		return nil, fmt.Errorf("could not cast todo item with id %s from map", id)
	}

	return &t, nil
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]todo.Todo, error) {
	var todos []todo.Todo

	s.m.Range(func(key, value any) bool {
		t, ok := value.(todo.Todo)
		if !ok {
			// TODO: better error propagation
			s.logger.Errorw("could not cast todo item from map", "todo_id", key, "raw_value", value)
			return false
		}

		todos = append(todos, t)
		return true
	})

	return todos, nil
}
