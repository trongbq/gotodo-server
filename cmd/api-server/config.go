package main

import "github.com/trongbq/gotodo-server/internal/config"

// appConfig contains all configuration for entire API server, from Server to other dependencies services.
type appConfig struct {
	env          string
	databaseURI  string
	authTokenKey string
}

// initConfig creates Config instance and must be called before any cofiguration values are used.
func initConfig() *appConfig {
	return &appConfig{
		env:          config.GetEnv("ENV", config.LocalEnv),
		databaseURI:  config.GetEnvStrict("DATABASE_URI"),
		authTokenKey: config.GetEnvStrict("AUTH_TOKEN_KEY"),
	}
}
