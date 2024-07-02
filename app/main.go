package main

import (
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {
	srv := server.New(4221)
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
