package main

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {

	router := http.NewRouter()

	router.Register("/echo/.*", http.EchoHandler)
	router.Register("^/$", http.RootHandler)

	srv := server.New(4221, router)

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
