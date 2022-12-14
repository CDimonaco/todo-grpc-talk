package todo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"go.uber.org/zap"
)

type Creator interface {
	CreateTodo(ctx context.Context, id, name, description string) (*Todo, error)
}

type AddTodo struct {
	creator Creator
	logger  *zap.SugaredLogger
}

func NewAddTodo(creator Creator, logger *zap.SugaredLogger) *AddTodo {
	l := logger.With("component", "todoCreator")

	return &AddTodo{
		logger:  l,
		creator: creator,
	}
}

func (c *AddTodo) Execute(ctx context.Context, name, description string) (*Todo, error) {
	if name == "" {
		return nil, fmt.Errorf("could not create new todo, missing name")
	}

	if description == "" {
		return nil, fmt.Errorf("could not create new todo, missing description")
	}

	newTodoId := uuid.NewString()

	todo, err := c.creator.CreateTodo(ctx, newTodoId, name, description)
	if err != nil {
		c.logger.Errorw("error during todo creation", "error", err)

		return nil, errors.Wrap(err, "could not create todo with todo creator")
	}

	return todo, nil
}
