package todo

import (
	"fmt"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

type Creator interface {
	CreateTodo(name, description string) (*Todo, error)
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

func (c *AddTodo) Execute(name, description string) (*Todo, error) {
	if name == "" {
		return nil, fmt.Errorf("could not create new todo, missing name")
	}

	if description == "" {
		return nil, fmt.Errorf("could not create new todo, missing description")
	}

	todo, err := c.creator.CreateTodo(name, description)
	if err != nil {
		c.logger.Errorw("error during todo creation", "error", err)

		return nil, errors.Wrap(err, "could not create todo with todo creator")
	}

	return todo, nil
}
