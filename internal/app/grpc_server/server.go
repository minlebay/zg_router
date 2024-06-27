package grpc_server

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	r "zg_router/internal/app/router"
	"zg_router/pkg/message_v1/router"
)

type Server struct {
	Logger *zap.Logger
	Config *Config
	router.UnimplementedMessageRouterServer
	Router     *r.Router
	GRPCServer *grpc.Server
	wg         sync.WaitGroup
}

func NewServer(logger *zap.Logger, config *Config, router *r.Router) *Server {
	return &Server{
		Logger:     logger,
		Config:     config,
		Router:     router,
		GRPCServer: grpc.NewServer(),
	}
}

func (s *Server) StartServer(ctx context.Context) {
	go func() {
		listener, err := net.Listen("tcp", s.Config.ListenAddress)
		if err != nil {
			s.Logger.Fatal(err.Error())
		}

		router.RegisterMessageRouterServer(s.GRPCServer, s)

		if err = s.GRPCServer.Serve(listener); err != nil {
			s.Logger.Fatal(err.Error())
		}

		s.Logger.Info("Server started at address " + s.Config.ListenAddress)
	}()
}

func (r *Server) StopServer(ctx context.Context) {
	r.wg.Wait()
	r.GRPCServer.Stop()
	r.Logger.Info("Server stopped")
}

func (s *Server) ReceiveMessage(ctx context.Context, m *router.Message) (*router.Response, error) {

	go func(m *router.Message) {
		s.wg.Add(1)
		defer s.wg.Done()

		err := s.Router.Route(ctx, m)
		if err != nil {
			s.Logger.Error(err.Error())
		}
	}(m)

	resp := router.Response{
		Success: true,
		Message: fmt.Sprintf("message received %v", time.Now()),
	}
	return &resp, nil
}
