package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Welcome...")
		fmt.Println("1. Send A Message")
		fmt.Println("2. Exit")
		scanner.Scan()
		opt := scanner.Text()

		if opt == "1" {
			SendMessageMenu()
		} else if opt == "2" {
			fmt.Println("GoodBye")
			break
		}
	}

}
func SendMessageMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	var Message string
	for {
		fmt.Printf("1. Please Enter The Message: ")
		scanner.Scan()
		Message = scanner.Text()

		if len(Message) <= 10 {
			fmt.Println("message Should Be More Than 10 Digits")
		} else if strings.Contains(Message, "kasar") {
			fmt.Println("No kata kasar Allowed")
		} else if strings.Compare(Message, "Hello World") == 0 {
			fmt.Printf("Your World Is %s\n, And This is not Allowed", Message)
		} else {
			break
		}
	}
	SendMessageToServer(Message)

}

func SendMessageToServer(message string) {
	ServerConnection, err := net.DialTimeout("tcp", "127.0.0.1:1234", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer ServerConnection.Close()

	err = binary.Write(ServerConnection, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}

	_, err = ServerConnection.Write([]byte(message))
	if err != nil {
		panic(err)
	}

}

func main() {
	menu()
}
