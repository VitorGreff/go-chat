package main

import (
	"flag"

	"chat/tcp"
)

var (
	ADDRESS_PORT    = ":8080"
	MESSAGE_CHANNEL = make(chan tcp.Message)
)

func main() {
	server := flag.Bool("s", false, "start server")
	flag.Parse()

	if *server {
		tcp.EchoServer(ADDRESS_PORT)
	} else {
		tcp.Client(ADDRESS_PORT, MESSAGE_CHANNEL)
	}
}
