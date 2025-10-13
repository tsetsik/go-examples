package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/tsetsik/go-examples/job-queue/internal/svc"
)

func main() {

	s, err := svc.NewService()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(
		gracefulStop,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	go func() {
		select {
		//nolint:errcheck
		case <-gracefulStop:
			cancel()
			s.Stop()
			os.Exit(0)
		case <-ctx.Done():
			return
		}
	}()

	if err := s.Start(ctx); err != nil {
		panic(err)
	}
}
