package main

import (
	"flag"
	"os"

	"chat/tcp"
)

var ADDRESS_PORT = ":8080"

func main() {
	message := os.Args[1]
	server := flag.Bool("s", false, "start server")
	flag.Parse()

	if *server {
		tcp.EchoServer(ADDRESS_PORT)
	} else {
		tcp.Client(ADDRESS_PORT, message)
	}
}
