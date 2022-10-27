package main

import (
	"context"
	"github.com/bastian-kurz/basic-rest-example/internal/logger"
	"github.com/bastian-kurz/basic-rest-example/internal/util"
	"github.com/bastian-kurz/basic-rest-example/server"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	os.Exit(start())
}

func start() int {
	defer func() {
		// If we cannot sync, there's probably something wrong with outputting logs,
		// so we probably cannot write using fmt.Println either. So just ignore the error.
		_ = logger.Log().Sync()
	}()

	host := util.GetStringOrDefault("HOST", "localhost")
	port := util.GetIntOrDefault("PORT", 8000)

	s := server.NewServer(server.Options{
		Host: host,
		Log:  logger.Log(),
		Port: port,
	})

	var eg errgroup.Group
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	eg.Go(func() error {
		<-ctx.Done()
		if err := s.Stop(); err != nil {
			logger.Log().Info("Error stopping server", zap.Error(err))
		}

		return nil
	})

	if err := s.Start(); err != nil {
		logger.Log().Info("Error starting server", zap.Error(err))
		return 1
	}

	if err := eg.Wait(); err != nil {
		return 1
	}

	return 0
}
