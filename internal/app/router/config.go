package router

import "go.uber.org/config"

type Config struct {
}

func NewRouterConfig(provider config.Provider) (*Config, error) {
	var config Config
	err := provider.Get("router").Populate(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
