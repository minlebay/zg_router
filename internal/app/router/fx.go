package router

import (
	"context"
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
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						go r.StartRouter()
						return nil
					},
					OnStop: func(context.Context) error {
						r.StopRouter()
						return nil
					},
				})
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("router")
		}),
	)
}
