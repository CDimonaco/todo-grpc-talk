package handlers

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GracefulShutdown(logger *zap.SugaredLogger, grpcServer *grpc.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	sig := <-signalChan
	logger.Infow("received signal, stopping grpc server", "signal", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	go func() {
		defer cancel()
		grpcServer.GracefulStop()
	}()
	<-ctx.Done()
}
