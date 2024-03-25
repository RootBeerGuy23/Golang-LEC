package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		ClientConnection, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go HandleServerConn(ClientConnection)
	}
}

func HandleServerConn(client net.Conn) {
	var size uint32
	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	byteMessage := make([]byte, size)
	_, err = client.Read(byteMessage)
	if err != nil {
		panic(err)
	}
	MessageAsli := string(byteMessage)
	fmt.Printf("Message Received: %s\n", MessageAsli)

}
