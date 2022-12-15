package client

import (
	"context"
	"fmt"

	"github.com/CDimonaco/todo-grpc-talk/internal/proto/grpc"
	grpc_lib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Completed   bool   `json:"completed"`
}

type TodoGrpc struct {
	baseURL    string
	gprcClient grpc.TodoServiceClient
	baseClient *grpc_lib.ClientConn
}

func NewTodoGrpcClient(baseURL string) (*TodoGrpc, error) {
	c, err := grpc_lib.Dial(
		baseURL,
		grpc_lib.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect to TodoGrpc instance: %s", baseURL)
	}

	mc := grpc.NewTodoServiceClient(c)

	m := TodoGrpc{
		baseURL:    baseURL,
		gprcClient: mc,
		baseClient: c,
	}

	return &m, nil
}

func (c *TodoGrpc) Close() error {
	return c.baseClient.Close()
}

func (c *TodoGrpc) AddTodo(ctx context.Context, name, description string) (*TodoResponse, error) {
	t, err := c.gprcClient.AddTodo(ctx, &grpc.AddTodoRequest{
		Title:       name,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	return &TodoResponse{
		ID:          t.GetResult().GetId(),
		Title:       t.GetResult().GetTitle(),
		Description: t.GetResult().GetDescription(),
		CreatedAt:   t.GetResult().GetCreatedAt(),
		Completed:   t.GetResult().GetCompleted(),
	}, nil
}

func (c *TodoGrpc) GetTodo(ctx context.Context, id string) (*TodoResponse, error) {
	t, err := c.gprcClient.GetTodo(ctx, &grpc.GetTodoRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &TodoResponse{
		ID:          t.GetResult().GetId(),
		Title:       t.GetResult().GetTitle(),
		Description: t.GetResult().GetDescription(),
		CreatedAt:   t.GetResult().GetCreatedAt(),
		Completed:   t.GetResult().GetCompleted(),
	}, nil
}

func (c *TodoGrpc) GetAllTodos(ctx context.Context) ([]TodoResponse, error) {
	list, err := c.gprcClient.ListTodos(ctx, &grpc.ListTodoRequest{})
	if err != nil {
		return nil, err
	}

	var todos []TodoResponse

	for _, t := range list.Todos {
		todos = append(todos, TodoResponse{
			ID:          t.GetId(),
			Title:       t.GetTitle(),
			Description: t.GetDescription(),
			CreatedAt:   t.GetCreatedAt(),
			Completed:   t.GetCompleted(),
		})
	}
	return todos, nil
}
