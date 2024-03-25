package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Proxy Is Running")
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
		go HandleProxyConnection(ClientConnection)
	}
}

func HandleProxyConnection(ClientConnection net.Conn) {
	ServerConnection, err := net.Dial("tcp", "127.0.0.1:4444")
	if err != nil {
		panic(ServerConnection)
	}
	defer ClientConnection.Close()
	defer ServerConnection.Close()

	go func() {
		_, err := io.Copy(ServerConnection, ClientConnection)
		if err != nil {
			panic(err)
		}
	}()

	_, err = io.Copy(ClientConnection, ServerConnection)
	if err != nil {
		panic(err)
	}
}
