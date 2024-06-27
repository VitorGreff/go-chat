package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client(port string) {
	address := fmt.Sprintf("127.0.0.1%s", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)

	for {
		var message string
		fmt.Print("\n> ")

		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error readling input")
			continue
		}
		sendMessage(conn, name, message)
	}
}

func sendMessage(conn net.Conn, name string, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error sending data to server: %s\n", err)
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

	// both values carry '\n' at the end
	fmt.Printf("> %s: %s> %s: %s", name, message, name, data[:n])
}
