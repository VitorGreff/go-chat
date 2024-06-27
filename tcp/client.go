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
	readUserInput(conn)
}

func readUserInput(conn net.Conn) string {
	defer conn.Close()
	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Printf("Welcome to the chat, %s!\n\n", name)

	reader := bufio.NewReader(os.Stdin)
	for {
		keyboardInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		message := fmt.Sprintf("%s: %s", name, keyboardInput)
		_, err = conn.Write([]byte(message))
		if err != nil {
			panic(err.Error())
		}
		fmt.Print(message)

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(buffer[:n]))
	}
}
