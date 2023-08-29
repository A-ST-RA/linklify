package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	
	AppName string `mapstructure:"APP_NAME"`
}

var AppConfig *Config

func InitializeConfig() error {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading env file: %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("error marshalling env file: %s", err)
	}
	
	return nil
}
