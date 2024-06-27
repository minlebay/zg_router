package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_router/internal/app/grpc_client"
	"zg_router/internal/app/grpc_server"
	"zg_router/internal/app/router"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			router.NewModule(),
			grpc_server.NewModule(),
			grpc_client.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
	)
}
