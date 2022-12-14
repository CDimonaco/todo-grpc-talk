package todo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type TodoGetter interface {
	GetTodoByID(ctx context.Context, id string) (*Todo, error)
}

type GetTodo struct {
	logger *zap.SugaredLogger
	getter TodoGetter
}

func NewGetTodo(
	logger *zap.SugaredLogger,
	getter TodoGetter,
) *GetTodo {
	l := logger.With("component", "getTodo")

	return &GetTodo{
		logger: l,
		getter: getter,
	}
}

func (g *GetTodo) Execute(ctx context.Context, id string) (*Todo, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id provided")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("the provided id %s is not an uuid", id)
	}

	t, err := g.getter.GetTodoByID(ctx, id)
	if err != nil {
		g.logger.Errorw("could not retrieve todo", "todo_id", id, "error", err)
		return nil, errors.Wrap(err, "could not retrieve todo")
	}

	return t, nil
}
