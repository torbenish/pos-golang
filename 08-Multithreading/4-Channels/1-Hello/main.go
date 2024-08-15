package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // Empty

	// Thread 2
	go func() {
		channel <- "Hello World!" // Full
	}()

	// Thread 1
	msg := <-channel
	fmt.Println(msg) // Channel will be empty
}
