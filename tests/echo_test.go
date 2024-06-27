package tests

import (
	"encoding/json"
	"fmt"
	"net"
	"testing"
	"time"

	"chat/tcp"
)

func TestNetworkCommunication(t *testing.T) {
	port := ":8080"

	go func() {
		tcp.EchoServer(port)
	}()
	time.Sleep(time.Second)

	address := fmt.Sprintf("127.0.0.1%s", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		t.Fatalf("Error connecting to server: %s", err)
	}
	defer conn.Close()

	message := tcp.Message{Source: "test", Content: "Hello, world!\n"}
	data, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Error marshalling message: %s", err)
	}

	_, err = conn.Write(append(data, '\n'))
	if err != nil {
		t.Fatalf("Error sending data to server: %s", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		t.Fatalf("Error receiving data from server: %s", err)
	}

	var receivedMessage tcp.Message
	err = json.Unmarshal(buffer[:n], &receivedMessage)
	if err != nil {
		t.Fatalf("Error unmarshalling received message: %s", err)
	}

	if receivedMessage.Source != message.Source || receivedMessage.Content != message.Content {
		t.Fatalf("Expected %v, but got %v", message, receivedMessage)
	}
}
