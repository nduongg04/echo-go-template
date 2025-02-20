package config

import (
	// "log"

	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	Env   string `mapstructure:"ENV"`
	DBUrl string `mapstructure:"DB_URL"`
  JWTSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.SetDefault("PORT", "8080")

	// if remove this, docker image will not load env
	viper.BindEnv("PORT")
	viper.BindEnv("ENV")
	viper.BindEnv("DB_URL")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Skipping .env file (Docker will inject env vars): %v", err)
	} else {
		log.Println("Loaded .env file")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
