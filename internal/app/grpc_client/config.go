package grpc_client

import "go.uber.org/config"

type Config struct {
	ProcessingServersList []string `yaml:"processing_servers_list"`
}

func NewClientConfig(provider config.Provider) (*Config, error) {
	var cfg Config

	err := provider.Get("processing").Populate(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
