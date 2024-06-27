package router

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {

	return fx.Module(
		"router",
		fx.Provide(
			NewRouterConfig,
			NewRouter,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r *Router) {
				lc.Append(fx.StartStopHook(r.StartRouter, r.StopRouter))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("router")
		}),
	)
}
