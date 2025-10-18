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
	"github.com/tsetsik/go-examples/job-queue/internal/infra"
	"github.com/tsetsik/go-examples/job-queue/internal/store"
	"github.com/tsetsik/go-examples/job-queue/internal/workers"

	"github.com/joho/godotenv"
)

type Service struct {
	cfg       config.Config
	infraSvc  infra.InfraService
	jobWorker core.JobWorker
}

func NewService() (*Service, error) {
	_, b, _, _ := runtime.Caller(0)
	if err := godotenv.Load(filepath.Dir(b) + "/../../.env"); err != nil {
		panic(err.Error())
	}

	host := getEnv[string]("HOST")
	infraPort := getEnv[int]("INFRA_PORT")
	port := getEnv[int]("PORT")
	numWorkers := getEnv[int]("JOB_WORKERS")

	cfg, err := config.LoadConfig(host, port, infraPort, numWorkers)
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

	s.infraSvc = infra.NewInfraService(s.cfg, logger)

	mux := mux.NewRouter()

	s.jobWorker = workers.NewJobWorker(5, logger)

	jobStore := store.NewCacheStore[core.Job]()
	jobSvc := core.NewJobQueueService(s.jobWorker, jobStore, logger)
	resolver := NewHttpResolver(ctx, logger, s.cfg, jobSvc)

	mux.HandleFunc("/status/{jobID}", resolver.Status).Methods(http.MethodGet)
	mux.HandleFunc("/submit", resolver.Submit).Methods(http.MethodPost)
	mux.HandleFunc("/health", resolver.Health).Methods(http.MethodGet)

	// pprof handlers
	mux.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)

	s.jobWorker.Start(jobSvc.ProcessedJob())

	// Start infra service (pprof, metrics, etc.)
	go s.infraSvc.Start()

	logger.Info("Starting server", slog.String("host", s.cfg.Host), slog.Int("port", s.cfg.Port))
	return http.ListenAndServe(s.cfg.Host+":"+strconv.Itoa(s.cfg.Port), mux)
}

func (s *Service) Stop() error {
	s.jobWorker.Stop()
	s.infraSvc.Stop()
	return nil
}

func getEnv[A any](key string) A {
	var zero A
	val, ok := os.LookupEnv(key)
	if !ok {
		return zero
	}

	var result any = zero
	switch any(zero).(type) {
	case string:
		result = val
	case int:
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return zero
		}
		result = parsed
	// Add more types as needed
	default:
		return zero
	}

	return result.(A)
}
