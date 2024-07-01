package router

import (
	"context"
	"go.uber.org/zap"
	"sync"
	"time"
	"zg_router/internal/app/grpc_client"
	"zg_router/pkg/message_v1"
)

type Router struct {
	Done   chan struct{}
	Logger *zap.Logger
	Config *Config
	Client *grpc_client.Client
	wg     sync.WaitGroup
}

func NewRouter(logger *zap.Logger, config *Config, client *grpc_client.Client) *Router {
	return &Router{
		Done:   make(chan struct{}),
		Logger: logger,
		Config: config,
		Client: client,
	}
}

func (r *Router) StartRouter(ctx context.Context) {
	go func() {
		for {
			select {
			case <-r.Done:
				return
			default:
				continue
			}
		}
	}()
}

func (r *Router) StopRouter(ctx context.Context) {
	r.wg.Wait()
	r.Done <- struct{}{}
	r.Logger.Info("Router stopped")
}

func (r *Router) Route(ctx context.Context, msg *message.Message) error {

	server := r.Client.GetLeastLoadedServer()
	if server == "" {
		r.Logger.Info("waiting for available server ...")
		time.Sleep(1 * time.Second)
	}

	go func() {
		r.wg.Add(1)
		defer r.wg.Done()
		r.Client.SendMessage(ctx, msg, server)
	}()

	return nil
}
