package grpc_server

import (
	cfg "go.uber.org/config"
)

type Config struct {
	ListenAddress string `yaml:"GRPC_SERVER_LISTEN_ADDRESS"`
}

func NewServerConfig(provider cfg.Provider) (*Config, error) {
	var config Config

	v := provider.Get("GRPC_SERVER_LISTEN_ADDRESS")
	var err = v.Populate(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
