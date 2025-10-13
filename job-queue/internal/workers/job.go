package workers

import (
	"log/slog"

	"github.com/tsetsik/go-examples/job-queue/internal/core"
)

type (
	jobWorker struct {
		logger     *slog.Logger
		numWorkers int
		jobs       chan core.Job
		results    chan core.JobProcessed
	}
)

func NewJobWorker(numWorkers int, logger *slog.Logger) core.JobWorker {
	return &jobWorker{
		numWorkers: numWorkers,
		logger:     logger,
		jobs:       make(chan core.Job),
		results:    make(chan core.JobProcessed),
	}
}

func (w *jobWorker) Start() {
	for i := 0; i < w.numWorkers; i++ {
		go w.processJob(w.jobs, w.results)
	}
}

func (w *jobWorker) Enqueue(job *core.Job) error {
	if err := job.Validate(); err != nil {
		w.logger.Error("invalid job", slog.String("error", err.Error()))
		return err
	}

	go func() {
		w.jobs <- *job
	}()

	return nil
}

func (w *jobWorker) processJob(jobs <-chan core.Job, results chan<- core.JobProcessed) {
	for job := range jobs {
		job.Status = core.StatusDone
		results <- core.JobProcessed{Job: job, Err: nil, Code: 0}
	}
}

func (w *jobWorker) Stop() {
	// Logic to gracefully stop workers
	close(w.jobs)
	close(w.results)
}
