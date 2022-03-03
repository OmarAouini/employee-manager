package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

//global config struct
var CONFIG *Configuration

func ConfigApp() {
	configLogger()
	configEnv()
}

//LOGGER

//config logger format
func configLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

//ENV

type Configuration struct {
	DatabaseName     string `env:"DATABASE_NAME"`
	DatabaseUser     string `env:"DATABASE_USERNAME"`
	DatabasePass     string `env:"DATABASE_PASSWORD"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabaseIdleConn int    `env:"DATABASE_IDLE_CONN"`
	DatabaseMaxConn  int    `env:"DATABASE_MAX_CONN"`
	IsProduction     bool   `env:"PRODUCTION"`
	JwtSecret        string `env:"JWT_SECRET"`
}

func loadDotEnv() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Error loading .env file")
	}
}

//load env variables
func configEnv() {
	loadDotEnv()
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)
	CONFIG = &cfg
}
