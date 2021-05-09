package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Listener struct {
			Port int
		}
		Logger struct {
			Version string
			Env     string
			Level   string
		}
	}
)

func NewAppConfig(configFile string) (*AppConfig, error) {
	config, err := loadConfig(configFile)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfig(configFile string) (*AppConfig, error) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	return &appConfig, errors.Wrap(err, "failed to get config")
}
