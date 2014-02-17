package main

import (
	"github.com/cgarvis/gestalt"
)

type Config struct {
	Driver   string
	Database string
	Host     string
	Port     string
	Key      string
}

func LoadConfig() Config {
	env := gestalt.New()

	config := Config{
		Driver:   env.String("citizens.driver", "postgres"),
		Database: env.String("citizens.database", "postgres://localhost/citizens"),
		Host:     env.String("citizens.host", "0.0.0.0"),
		Port:     env.String("cizizens.port", "3000"),
	}

	return config
}
