package handlers

import (
	"context"
	"time"

	"github.com/CDimonaco/todo-grpc-talk/internal/proto/grpc"
	"github.com/CDimonaco/todo-grpc-talk/internal/proto/stubs"
	"github.com/CDimonaco/todo-grpc-talk/internal/todo"
	"go.uber.org/zap"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type GrpcHandler struct {
	grpc.UnimplementedTodoServiceServer
	logger     *zap.SugaredLogger
	addTodo    *todo.AddTodo
	getTodo    *todo.GetTodo
	getAllTodo *todo.GetAllTodo
}

var _ grpc.TodoServiceServer = (*GrpcHandler)(nil)

func NewGrpcHandler(
	logger *zap.SugaredLogger,
	addTodo *todo.AddTodo,
	getTodo *todo.GetTodo,
	getAllTodo *todo.GetAllTodo,
) *GrpcHandler {
	l := logger.With("component", "grpcHandler")

	return &GrpcHandler{
		logger:     l,
		addTodo:    addTodo,
		getTodo:    getTodo,
		getAllTodo: getAllTodo,
	}
}

func (h *GrpcHandler) AddTodo(ctx context.Context, req *grpc.AddTodoRequest) (*grpc.AddTodoResponse, error) {
	newTodo, err := h.addTodo.Execute(ctx, req.Title, req.Description)
	// TODO: propagate specific error from the usecase to distinguish bad requests
	// from generic errors
	if err != nil {
		h.logger.Error("error during todo add", "err", err)
		return nil, status.Errorf(codes.Internal, "error during todo creation %s", err)
	}

	return &grpc.AddTodoResponse{
		Result: mapTodoEntityToProtoTodo(newTodo),
	}, nil
}

func (h *GrpcHandler) GetTodo(ctx context.Context, in *grpc.GetTodoRequest) (*grpc.GetTodoResponse, error) {
	t, err := h.getTodo.Execute(ctx, in.GetId())
	if err != nil {
		h.logger.Error("error during todo get", "err", err)
		return nil, status.Errorf(codes.Internal, "error during todo get %s", err)
	}

	return &grpc.GetTodoResponse{
		Result: mapTodoEntityToProtoTodo(t),
	}, nil
}

func (h *GrpcHandler) ListTodos(ctx context.Context, req *grpc.ListTodoRequest) (*grpc.ListTodosResponse, error) {
	todos, err := h.getAllTodo.Execute(ctx)
	if err != nil {
		h.logger.Error("error during todo get all", "err", err)
		return nil, status.Errorf(codes.Internal, "error during todo get all %s", err)
	}
	var todosResponse []*stubs.Todo

	for _, t := range todos {
		todoStub := mapTodoEntityToProtoTodo(&t)
		todosResponse = append(todosResponse, todoStub)
	}

	return &grpc.ListTodosResponse{
		Todos: todosResponse,
	}, nil
}

func mapTodoEntityToProtoTodo(entity *todo.Todo) *stubs.Todo {
	return &stubs.Todo{
		Id:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt.Format(time.RFC3339),
		Completed:   entity.Completed,
	}
}
