package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	StoragePath string `mapstructure:"storage_path"`
	LogPath     string `mapstructure:"log_path"`
	TimeZone    string `mapstructure:"timezone"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("storage_path", "./data")
	viper.SetDefault("log_path", "./logs")
	viper.SetDefault("timezone", "UTC")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.task-tracker")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
