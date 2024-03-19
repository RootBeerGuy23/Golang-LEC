package services

import "fmt"

func PrintNumber() {
	for i := 1; i < 10; i++ {
		fmt.Printf("This is the number: %d\n", i)
	}
}
