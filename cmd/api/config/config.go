package config

import (
	"github.com/trongbq/gotodo-server/internal/config"
)

// Config contains all configuration for entire API server, from Server to other dependencies services.
type Config struct {
	Env         string
	DatabaseURI string
}

// Init creates Config instance and must be called before any cofiguration values are used.
func Init() *Config {
	return &Config{
		Env:         config.GetEnv("ENV", config.LocalEnv),
		DatabaseURI: config.GetEnvStrict("DATABASE_URI"),
	}
}
