package grpc_client

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {

	return fx.Module(
		"client",
		fx.Provide(
			NewClientConfig,
			NewClient,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, g *Client) {
				lc.Append(fx.StartStopHook(g.StartClient, g.StopClient))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("client")
		}),
	)
}
