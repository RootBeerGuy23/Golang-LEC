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
		fmt.Println("1. Send Message To Server")
		fmt.Println("2. Exit")
		scanner.Scan()
		opt := scanner.Text()
		if opt == "1" {
			MessageMenu()
		} else if opt == "2" {
			fmt.Println("bye!")
			break
		}
	}
}

func MessageMenu() {
	var message string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please Input Your Message:")
		scanner.Scan()
		message = scanner.Text()
		if len(message) < 10 {
			fmt.Print("Message Should Be More Than 10 Character")
		} else if strings.Contains(message, "kasar") {
			fmt.Println("No Bad Words Allowed")
		} else if strings.Compare(message, "helloworld123") == 0 {
			fmt.Printf("Your Message Is %s, And Its Not Allowed", message)
		} else {
			break
		}
	}
	KirimKeServer(message)
}

func KirimKeServer(message string) {
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
