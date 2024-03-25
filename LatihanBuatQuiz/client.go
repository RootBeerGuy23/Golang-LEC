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

func Menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1.Send Message To Server")
		fmt.Println("2.Exit")
		scanner.Scan()
		choose := scanner.Text()
		if choose == "1" {
			MessageMenu()
		} else if choose == "2" {
			fmt.Print("Bye")
			break
		}
	}
}

func MessageMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	var Message string
	for {
		fmt.Print("Insert Your Message: ")
		scanner.Scan()
		Message = scanner.Text()
		if len(Message) < 10 {
			fmt.Print("Message Should Be More Than 10 Char")
		} else if strings.Contains(Message, "kasar") {
			fmt.Printf("No Bad Words Allowed")
		} else if strings.Compare(Message, "hello world") == 0 {
			fmt.Printf("Your Message is : %s\n & this is not allowed", Message)
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

	BytReply := make([]byte, size)
	_, err = ServerConnection.Read(BytReply)
	if err != nil {
		panic(err)
	}

	RealReply := string(BytReply)
	fmt.Printf("Replied From Server %s\n", RealReply)

}

func main() {
	Menu()
}
