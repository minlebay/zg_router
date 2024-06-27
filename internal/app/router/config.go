package router

import "go.uber.org/config"

type Config struct {
}

func NewRouterConfig(provider config.Provider) (*Config, error) {

	var config Config
	return &config, nil

	err := provider.Get("processing").Populate(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
