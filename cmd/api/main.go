package main

import (
	"github.com/trongbq/gotodo-server/internal/api"
)

func main() {
	sconf := api.ServerConfig{}
	server := api.NewServer(sconf)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
