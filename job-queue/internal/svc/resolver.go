package svc

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tsetsik/go-examples/job-queue/internal/config"
	"github.com/tsetsik/go-examples/job-queue/internal/core"
)

type (
	HttpResolver interface {
		Submit(w http.ResponseWriter, r *http.Request)
		Status(w http.ResponseWriter, r *http.Request)
	}

	httpResolver struct {
		ctx    context.Context
		logger *slog.Logger
		config config.Config
		jobSvc core.JobQueueService
	}
)

func NewHttpResolver(ctx context.Context, logger *slog.Logger, cfg config.Config, jobSvc core.JobQueueService) HttpResolver {
	return &httpResolver{
		ctx:    ctx,
		logger: logger,
		config: cfg,
		jobSvc: jobSvc,
	}
}

func (s *httpResolver) Submit(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	job := &core.Job{}
	decoder.Decode(job)

	err := s.jobSvc.Enqueue(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (s *httpResolver) Status(w http.ResponseWriter, r *http.Request) {
	// Implementation for checking job status
}
