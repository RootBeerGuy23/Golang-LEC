package main

import (
	"io"
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
		go HandleProxyConnection(ClientConnection)
	}
}

func HandleProxyConnection(ClientConnection net.Conn) {
	ServerConnection, err := net.Dial("tcp", "127.0.0.1:4444")
	if err != nil {
		panic(err)
	}
	defer ServerConnection.Close()
	defer ClientConnection.Close()

	go func() {
		_, err := io.Copy(ServerConnection, ClientConnection)
		if err != nil {
			panic(err)
		}
	}()

	io.Copy(ClientConnection, ServerConnection)
	if err != nil {
		panic(err)
	}
}
