package main

import (
	"fmt"
	"net"
	"os"
)

// Uncomment this block to pass the first stage
// "net"
// "os"

func main() {

	port := 4221

	fmt.Println("Starting http server in port:", port)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 4221))
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
