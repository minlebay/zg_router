package app

import (
	"go.uber.org/config"
	"go.uber.org/fx"
	"os"
)

type Config struct {
	Name string `yaml:"name"`
}

type ResultConfig struct {
	fx.Out
	Provider config.Provider
	Config   Config
}

func NewConfig() (ResultConfig, error) {

	cfg := config.Static(os.Environ())
	//loader, err := config.NewYAML(config.Static("config.yaml"))
	loader, err := config.NewYAML(cfg)
	if err != nil {
		return ResultConfig{}, err
	}

	config := Config{
		Name: "default",
	}

	if err = loader.Get("app").Populate(&config); err != nil {
		return ResultConfig{}, err
	}

	return ResultConfig{
		Provider: loader,
		Config:   config,
	}, nil

}
