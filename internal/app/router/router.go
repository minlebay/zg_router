package router

import "go.uber.org/zap"

type Router struct {
	Done   chan struct{}
	Logger *zap.Logger
	Config *Config
}

func NewRouter(logger *zap.Logger, config *Config) *Router {
	return &Router{
		Done:   make(chan struct{}),
		Logger: logger,
		Config: config,
	}
}

func (r *Router) StartRouter() {
	r.Logger.Info("Router started")
}

func (r *Router) StopRouter() {
	r.Logger.Info("Router stopped")
	r.Done <- struct{}{}
}
