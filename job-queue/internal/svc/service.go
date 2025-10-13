package svc

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tsetsik/go-examples/job-queue/internal/config"
	"github.com/tsetsik/go-examples/job-queue/internal/core"
	"github.com/tsetsik/go-examples/job-queue/internal/workers"

	"github.com/joho/godotenv"
)

type Service struct {
	cfg config.Config
}

func NewService() (*Service, error) {
	_, b, _, _ := runtime.Caller(0)
	if err := godotenv.Load(filepath.Dir(b) + "/../../.env"); err != nil {
		panic(err.Error())
	}

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	numWorkers, err := strconv.Atoi(os.Getenv("JOB_WORKERS"))
	if err != nil {
		return nil, fmt.Errorf("invalid number job workers: %w", err)
	}

	cfg, err := config.LoadConfig(host, port, numWorkers)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &Service{
		cfg: *cfg,
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	// Validate the config
	if err := s.cfg.Validate(); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	mux := mux.NewRouter()
	jobWorker := workers.NewJobWorker(5, logger)

	jobSvc := core.NewJobQueueService(jobWorker)
	resolver := NewHttpResolver(ctx, logger, s.cfg, jobSvc)

	mux.HandleFunc("/status{job_id}", resolver.Status).Methods(http.MethodGet)
	mux.HandleFunc("/submit", resolver.Submit).Methods(http.MethodPost)

	jobWorker.Start()

	logger.Info("Starting server", slog.String("host", s.cfg.Host), slog.Int("port", s.cfg.Port))
	return http.ListenAndServe(s.cfg.Host+":"+strconv.Itoa(s.cfg.Port), mux)
}

func (s *Service) Stop() error {
	// Implement any necessary cleanup logic here
	return nil
}
