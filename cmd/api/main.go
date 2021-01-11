package main

import (
	"log"
	"net/http"

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
	log.Fatal(http.ListenAndServe(":8080", server))
}
