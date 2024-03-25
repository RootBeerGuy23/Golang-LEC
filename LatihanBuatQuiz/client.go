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
		fmt.Println("Welcome")
		fmt.Println("1. Write Message To Server")
		fmt.Println("2. Exit")
		scanner.Scan()
		option := scanner.Text()
		if option == "1" {
			MessageMenu()
		} else if option == "2" {
			break
		}
	}
}

func MessageMenu() {
	var Message string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please Insert The Message: ")
		scanner.Scan()
		Message = scanner.Text()
		if len(Message) < 10 {
			fmt.Println("Should Me More Than 10 Characters")
		} else if strings.Contains(Message, "kasar") {
			fmt.Printf("No Bad Words Allowed")
		} else if strings.Compare(Message, "Hello World") == 0 {
			fmt.Printf("Your Message Is %s\n, This Is Not Allowed", Message)
		} else {
			break
		}
	}
	SendMessageToServer(Message)
}

func SendMessageToServer(Message string) {
	ServerConnection, err := net.DialTimeout("tcp", "127.0.0.1:1234", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer ServerConnection.Close()

	err = binary.Write(ServerConnection, binary.LittleEndian, uint32(len(Message)))
	if err != nil {
		panic(err)
	}
	_, err = ServerConnection.Write([]byte(Message))
	if err != nil {
		panic(err)
	}

	var size uint32
	err = binary.Read(ServerConnection, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	bytreplymsg := make([]byte, size)
	_, err = ServerConnection.Read(bytreplymsg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Replied From Server: %s\n", string(bytreplymsg))

}

func main() {
	menu()
}
