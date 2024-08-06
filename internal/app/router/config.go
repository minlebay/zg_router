package router

import "go.uber.org/config"

type Config struct {
}

func NewRouterConfig(provider config.Provider) (*Config, error) {
	var config Config
	return &config, nil
}
