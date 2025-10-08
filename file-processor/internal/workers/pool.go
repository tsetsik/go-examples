package workers

import (
	"log/slog"

	"github.com/tsetsik/go-examples/file-processor/internal/core"
)

type (
	WorkerPool interface {
		ProcessFileWorker(files <-chan string, results chan<- core.FileProcessed)
	}

	workerPool struct {
		logger *slog.Logger
	}
)

func NewWorkerPool(logger *slog.Logger) WorkerPool {
	return &workerPool{
		logger: logger,
	}
}
