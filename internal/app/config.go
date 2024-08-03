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
	yamlProvider, err := config.NewYAML(
		config.File("config.yaml"),
		config.Expand(os.LookupEnv),
	)
	if err != nil {
		return ResultConfig{}, err
	}

	var yamlConfig map[interface{}]interface{}
	if err = yamlProvider.Get(config.Root).Populate(&yamlConfig); err != nil {
		return ResultConfig{}, err
	}

	config := Config{
		Name: "default",
	}

	if err = yamlProvider.Get("app").Populate(&config); err != nil {
		return ResultConfig{}, err
	}

	return ResultConfig{
		Provider: yamlProvider,
		Config:   config,
	}, nil
}
