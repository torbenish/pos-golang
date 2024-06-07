package main

import "fmt"

func main() {
	defer fmt.Println("First Line")
	fmt.Println("Seconf Line")
	fmt.Println("Third Line")
}
