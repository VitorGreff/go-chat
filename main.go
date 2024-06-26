package main

import (
	"chat/tcp"
)

var ADDRESS_PORT = ":8080"

func main() {
	tcp.EchoServer(ADDRESS_PORT)
}
