package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	appConfig *AppConfig
)

type AppConfig struct {
	AppPort    int    `mapstructure:"APP_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSchema   string `mapstructure:"DB_SCHEMA"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     int    `mapstructure:"DB_PORT"`
	SecretKey  string `mapstructure:"SECRET_KEY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config *AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config not found, use env variable, current path = ", path)
			// Config file not found; ignore error if desired
		}
	}
	err = viper.Unmarshal(&config)
	return
}

// getConfig Initiatilize config
func getConfig() *AppConfig {
	if appConfig == nil {
		var err error

		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)
		appConfig, err = LoadConfig(basepath)
		if err != nil {
			log.Fatal("cannot load config:", err)
		}
	}

	return appConfig
}

func GetConfig() *AppConfig {
	return getConfig()
}
