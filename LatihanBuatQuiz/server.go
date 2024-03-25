package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Server Is Listening to 127.0.0.1:4444")
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

	bytmsg := make([]byte, size)
	_, err = ClientConnection.Read(bytmsg)

	Realmsg := string(bytmsg)
	fmt.Printf("Received: %s\n", Realmsg)

	var reply string
	if strings.HasSuffix(Realmsg, ".zip") {
		reply = "Received ZIP File"
	} else if strings.Contains(Realmsg, ".") {
		reply = "Only ZIP File Can Be Received"
	} else {
		reply = "message Has Been Received"
	}

	binary.Write(ClientConnection, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}

	_, err = ClientConnection.Write([]byte(reply))
	if err != nil {
		panic(err)
	}

}
