package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {

	var filesPath string

	if slices.Contains(os.Args, "--directory") && len(os.Args) > 2 {
		filesPath = os.Args[2]
	}

	fmt.Println("Configured files directory:", filesPath)

	router := http.NewRouter()

	router.Register("/echo/.*", "GET", http.EchoHandler)
	router.Register("/user-agent", "GET", http.UserAgentHandler)
	router.Register("/$", "GET", http.RootHandler)
	router.Register("/files/.*", "GET", http.FileHandler(filesPath))
	router.Register("/files/.*", "POST", http.FilePostHandler(filesPath))

	srv := server.New(4221, router)

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
