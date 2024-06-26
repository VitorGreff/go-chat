package tcp

import (
	"fmt"
	"io"
	"net"
)

func EchoServer(port string) {
	server, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	fmt.Printf("Server running on port: %s", port)

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s", err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			io.Copy(conn, conn)
		}(conn)
	}
}
