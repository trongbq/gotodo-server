package main

import (
	"github.com/trongbq/gotodo-server/cmd/api/config"
	"github.com/trongbq/gotodo-server/internal/api"
)

func main() {
	conf := config.Init()

	server := api.NewServer(
		api.ServerConfig{
			Env: conf.Env,
		},
	)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
