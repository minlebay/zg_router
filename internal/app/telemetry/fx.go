package telemetry

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {

	return fx.Module(
		"metrics",
		fx.Provide(
			NewMetricsConfig,
			NewMetrics,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, m *Metrics) {
				lc.Append(fx.StartStopHook(m.StartMetricsServer, m.StopMetricsServer))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("router_metrics")
		}),
	)
}
