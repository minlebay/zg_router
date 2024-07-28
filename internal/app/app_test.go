package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
	"zg_router/internal/app/grpc_client"
	"zg_router/internal/app/grpc_server"
	"zg_router/internal/app/log"
	"zg_router/internal/app/router"
	"zg_router/internal/app/telemetry"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
		fx.Options(
			router.NewModule(),
			grpc_server.NewModule(),
			grpc_client.NewModule(),
			log.NewModule(),
			telemetry.NewModule(),
		),
		fx.Provide(
			NewConfig,
		))
	require.NoError(t, err)
}
