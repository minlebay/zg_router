package grpc_server

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	"zg_router/internal/app/router"
	"zg_router/pkg/message_v1"
)

type Server struct {
	Logger *zap.Logger
	Config *Config
	message.UnimplementedMessageRouterServer
	Router     *router.Router
	GRPCServer *grpc.Server
	wg         sync.WaitGroup
}

func NewServer(logger *zap.Logger, config *Config, router *router.Router) *Server {
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

		message.RegisterMessageRouterServer(s.GRPCServer, s)

		if err = s.GRPCServer.Serve(listener); err != nil {
			s.Logger.Fatal(err.Error())
		}
	}()
}

func (s *Server) StopServer(ctx context.Context) {
	s.wg.Wait()
	s.GRPCServer.Stop()
	s.Logger.Info("Server stopped")
}

func (s *Server) ReceiveMessage(ctx context.Context, m *message.Message) (*message.Response, error) {

	go func(m *message.Message) {
		s.wg.Add(1)
		defer s.wg.Done()

		err := s.Router.Route(ctx, m)
		if err != nil {
			s.Logger.Error(err.Error())
		}
	}(m)

	resp := message.Response{
		Success: true,
		Message: fmt.Sprintf("message received %v", time.Now()),
	}
	return &resp, nil
}
