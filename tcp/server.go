package tcp

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func EchoServer(port string) {
	server, err := net.Listen("tcp", port)
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()
	fmt.Printf("Server running on port: %s\n", port)

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err.Error())
				break
			}
		}

		_, err = io.Copy(conn, bytes.NewReader(buffer[:n]))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
