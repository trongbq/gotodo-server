package main

import (
	"github.com/trongbq/gotodo-server/cmd/api/config"
	"github.com/trongbq/gotodo-server/internal/api"
	"github.com/trongbq/gotodo-server/internal/database"
)

func main() {
	conf := config.Init()

	db, err := database.New(conf.DatabaseURI)
	if err != nil {
		panic(err)
	}
	server := api.NewServer(
		api.ServerConfig{
			Env: conf.Env,
		},
		db,
	)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
