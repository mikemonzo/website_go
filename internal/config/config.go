package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./web/config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg := &Config{
		HTTPPort: viper.GetString("http.port"),
	}

	return cfg
}
