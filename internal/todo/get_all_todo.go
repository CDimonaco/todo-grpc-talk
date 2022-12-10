package todo

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AllTodoGetter interface {
	GetAllTodos(ctx context.Context) ([]Todo, error)
}

type GetAllTodo struct {
	logger *zap.SugaredLogger
	getter AllTodoGetter
}

func NewGetAllTodo(
	logger *zap.SugaredLogger,
	getter AllTodoGetter,
) *GetAllTodo {
	l := logger.With("component", "getAllTodo")

	return &GetAllTodo{
		logger: l,
		getter: getter,
	}
}

func (g *GetAllTodo) Execute(ctx context.Context) ([]Todo, error) {
	t, err := g.getter.GetAllTodos(ctx)
	if err != nil {
		g.logger.Errorw("could not get all todos", "error", err)
		return nil, errors.Wrap(err, "could not get all todos from all todo getter")
	}

	return t, nil
}
