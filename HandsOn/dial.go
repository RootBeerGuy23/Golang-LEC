package main

import "net"

func main() {
	conn, err := net.Dial("tcp","127.0.0.1:1234")
	if err != nil {
		return err
	}


}