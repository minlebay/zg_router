package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"testing"
	"zg_router/internal/app/grpc_client"
	"zg_router/internal/app/grpc_server"
	"zg_router/internal/app/router"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
		fx.Options(
			router.NewModule(),
			grpc_server.NewModule(),
			grpc_client.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		))
	require.NoError(t, err)
}
