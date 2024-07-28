package app

import (
	"go.uber.org/fx"
	"zg_router/internal/app/grpc_client"
	"zg_router/internal/app/grpc_server"
	"zg_router/internal/app/log"
	"zg_router/internal/app/router"
	"zg_router/internal/app/telemetry"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			router.NewModule(),
			grpc_server.NewModule(),
			grpc_client.NewModule(),
			log.NewModule(),
			telemetry.NewModule(),
		),
		fx.Provide(
			NewConfig,
		),
	)
}
