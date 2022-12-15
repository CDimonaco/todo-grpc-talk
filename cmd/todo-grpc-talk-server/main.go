package main

import (
	"net"
	"time"

	"github.com/CDimonaco/todo-grpc-talk/internal/handlers"
	"github.com/CDimonaco/todo-grpc-talk/internal/persistence"
	"github.com/CDimonaco/todo-grpc-talk/internal/proto/grpc"
	"github.com/CDimonaco/todo-grpc-talk/internal/todo"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	grpc_server "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var Version = "development"
var BuildDate = time.Now().Format("Mon Jan 2 15:04:05")

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	sugaredLogger := logger.Sugar()

	todoService := persistence.NewTodoService(sugaredLogger)
	addTodo := todo.NewAddTodo(todoService, sugaredLogger)
	getTodo := todo.NewGetTodo(sugaredLogger, todoService)
	getAllTodos := todo.NewGetAllTodo(sugaredLogger, todoService)

	sugaredLogger.Infow(
		"starting todo-grpc-talk-server",
		"version",
		Version,
		"build_date",
		BuildDate,
	)

	grpcHandlers := handlers.NewGrpcHandler(
		sugaredLogger,
		addTodo,
		getTodo,
		getAllTodos,
	)
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %v", err)
	}

	s := grpc_server.NewServer(
		grpc_server.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	grpc.RegisterTodoServiceServer(s, grpcHandlers)
	reflection.Register(s)
	// run signal handling goroutine
	go handlers.GracefulShutdown(sugaredLogger, s)

	sugaredLogger.Infof("listening on %s", "0.0.0.0:9090")
	if err := s.Serve(lis); err != nil {
		sugaredLogger.Fatalf("failed to serve: %v", err)
	}
}
