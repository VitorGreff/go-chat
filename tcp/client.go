package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client(port string, messageChannel chan Message) {
	address := fmt.Sprintf("127.0.0.1%s", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	go readUserInput(messageChannel)
	for message := range messageChannel {
		fmt.Printf("%s: %s> %s: %s\n", message.Source, message.Content, message.Source, message.Content)
	}
}

func readUserInput(messageChannel chan Message) {
	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		messageContent, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}
		messageChannel <- Message{Source: name, Content: messageContent}
	}
}
