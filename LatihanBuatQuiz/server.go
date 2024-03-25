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
		go HandleServerConnection(ClientConnection)
	}
}

func HandleServerConnection(Client net.Conn) {
	var size uint32
	err := binary.Read(Client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	BytMsg := make([]byte, size)
	_, err = Client.Read(BytMsg)
	if err != nil {
		panic(err)
	}
	StrMsg := string(BytMsg)
	fmt.Printf("Message Received: %s\n", StrMsg)
}
