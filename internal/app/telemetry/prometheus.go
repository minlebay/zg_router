package telemetry

import (
	"context"
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
)

type Metrics struct {
	Done           chan struct{}
	Logger         *zap.Logger
	Config         *Config
	RequestCounter prometheus.Counter
	server         *http.Server
}

func NewMetrics(logger *zap.Logger, config *Config) *Metrics {
	requestCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "zg_router_grpc_calls_total",
		Help: "Total number of GRPC calls",
	})

	prometheus.MustRegister(requestCounter)

	return &Metrics{
		Done:           make(chan struct{}),
		Logger:         logger,
		Config:         config,
		RequestCounter: requestCounter,
	}
}

func (m *Metrics) StartMetricsServer() {
	m.server = &http.Server{Addr: m.Config.Url, Handler: nil}
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		m.Logger.Info("Starting Prometheus metrics server on " + m.Config.Url)
		if err := m.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			m.Logger.Fatal("Failed to start Prometheus metrics server", zap.Error(err))
		}
	}()
}

func (m *Metrics) StopMetricsServer(ctx context.Context) {
	if err := m.server.Shutdown(ctx); err != nil {
		m.Logger.Fatal("Failed to gracefully shutdown the server", zap.Error(err))
	}
	m.Done <- struct{}{}
}

func (m *Metrics) IncrementRequestCounter() {
	m.RequestCounter.Inc()
}
