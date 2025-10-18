package infra

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tsetsik/go-examples/job-queue/internal/config"
)

type (
	InfraService interface {
		Start()
		Stop()
	}

	infraService struct {
		cfg    config.Config
		logger *slog.Logger
	}
)

func NewInfraService(cfg config.Config, logger *slog.Logger) InfraService {
	return &infraService{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *infraService) Start() {
	router := mux.NewRouter()
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	router.Handle("/metrics", promhttp.Handler())

	s.logger.Info("Starting infra service", "host", s.cfg.Host, "port", s.cfg.InfraPort)

	http.ListenAndServe(s.cfg.Host+":"+strconv.Itoa(s.cfg.InfraPort), router)
}

func (s *infraService) Stop() {
	// Implement graceful shutdown if needed
}
