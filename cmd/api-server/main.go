package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal/api"
	"github.com/trongbq/gotodo-server/internal/config"
	"github.com/trongbq/gotodo-server/internal/database"
)

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
func main() {
	conf := initConfig()
	initLogging(conf.env)

	db, err := database.New(conf.databaseURI)
	if err != nil {
		logrus.Panic(err)
	}
	server := api.NewServer(
		api.ServerConfig{
			Env:          conf.env,
			AuthTokenKey: conf.authTokenKey,
		},
		db,
	)
	logrus.WithFields(logrus.Fields{
		"host": "http://127.0.0.1",
		"port": "8080",
	}).Info("starting the HTTP server")
	logrus.Panic(http.ListenAndServe(":8080", server))
}

func initLogging(env string) {
	if env == config.LocalEnv || env == config.TestEnv {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Prefer using structured log
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of stderr
	logrus.SetOutput(os.Stdout)
}
