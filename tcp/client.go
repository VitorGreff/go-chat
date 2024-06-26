package tcp

import (
	"fmt"
	"net"
)

func Client(port string, message string) {
	address := fmt.Sprintf("127.0.0.1%s", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Error connecting to server: %s", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error sending data to server: %s", err)
		return
	}

	// allocate 1 gigabyte of memory to store the data
	data := make([]byte, 1024)
	// n -> number of bytes read (offset)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Printf("Error receiving data from server: %s", err)
		return
	}

	fmt.Printf("Received data from server: %s", data[:n])
}
