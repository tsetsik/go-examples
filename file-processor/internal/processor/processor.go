package processor

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/tsetsik/go-examples/file-processor/internal/core"
	w "github.com/tsetsik/go-examples/file-processor/internal/workers"
)

type (
	FilesProcessor interface {
		Process(filesPath []string) []core.FileProcessed
	}

	filesProcessor struct {
		wg         sync.WaitGroup
		m          sync.Mutex
		logger     *slog.Logger
		workers    int
		results    chan core.FileProcessed
		jobs       chan string
		output     chan core.FileProcessed
		workerPool w.WorkerPool
	}
)

func NewFilesProcessor(logger *slog.Logger, workers int) FilesProcessor {
	fp := &filesProcessor{
		workers:    workers,
		results:    make(chan core.FileProcessed),
		jobs:       make(chan string),
		logger:     logger,
		workerPool: w.NewWorkerPool(logger),
	}

	fp.spawnWorkers()

	return fp
}

func (fp *filesProcessor) spawnWorkers() {
	// Spawn file workers
	for i := 0; i < fp.workers; i++ {
		fp.logger.Info(fmt.Sprintf("Spawning worker %d", i+1))
		go fp.workerPool.ProcessFileWorker(fp.jobs, fp.results)
	}
}

func (fp *filesProcessor) Process(filesPath []string) []core.FileProcessed {
	fp.m.Lock()
	defer fp.m.Unlock()

	fp.output = make(chan core.FileProcessed, len(filesPath))

	fp.wg.Add(len(filesPath))
	for _, path := range filesPath {
		go func(p string) {
			fp.logger.Info("Sending job to worker", slog.String("file", p))
			fp.jobs <- p
		}(path)
	}

	go func() {
		for f := range fp.results {
			fp.logger.Info("Received result from worker", slog.Any("file", f))
			fp.wg.Done()

			fp.output <- f
		}
	}()

	fp.wg.Wait()

	// All jobs are done, shutdown the channels
	fp.shutdown()

	output := make([]core.FileProcessed, 0, len(filesPath))
	for res := range fp.output {
		output = append(output, res)
	}

	return output
}

func (fp *filesProcessor) shutdown() {
	fp.logger.Info("Shutting down workers")
	close(fp.jobs)
	close(fp.results)
	close(fp.output)
}
