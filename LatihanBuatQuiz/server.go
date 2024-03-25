package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:4444")
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

func HandleServerConnection(ClientConnection net.Conn) {
	var size uint32

	err := binary.Read(ClientConnection, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	BytMsg := make([]byte, size)
	ClientConnection.Read(BytMsg)

	RealMessage := string(BytMsg)
	fmt.Printf("Received: %s\n", RealMessage)

	var reply string
	if strings.HasSuffix(RealMessage, ".zip") {
		reply = "ZIP File Has Been Received"
	} else if strings.Contains(RealMessage, ".") {
		reply = "File Must Be ZIP"
	} else {
		reply = "Message Has Been Received"
	}
	err = binary.Write(ClientConnection, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}
	_, err = ClientConnection.Write([]byte(reply))
	if err != nil {
		panic(err)
	}

}
