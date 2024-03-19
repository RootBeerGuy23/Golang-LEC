package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("ini goroutine 1")

	go fmt.Println("ini goroutine 2")
	
	go fmt.Println("ini goroutine 3")

	time.Sleep(1 * time.Second)
}