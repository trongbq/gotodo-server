package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal/api"
	"github.com/trongbq/gotodo-server/internal/config"
	"github.com/trongbq/gotodo-server/internal/database"
)

func main() {
	conf := initConfig()
	initLogging(conf.env)

	db, err := database.New(conf.databaseURI)
	if err != nil {
		panic(err)
	}
	server := api.NewServer(
		api.ServerConfig{
			Env:          conf.env,
			AuthTokenKey: conf.authTokenKey,
		},
		db,
	)
	log.WithFields(log.Fields{
		"host": "http://127.0.0.1",
		"port": "8080",
	}).Info("starting the HTTP server")
	log.Panic(http.ListenAndServe(":8080", server))
}

func initLogging(env string) {
	if env == config.LocalEnv || env == config.TestEnv {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Prefer using structured log
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of stderr
	log.SetOutput(os.Stdout)
}
