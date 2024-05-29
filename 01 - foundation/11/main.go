package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	john := Client{
		Name:   "John",
		Age:    30,
		Active: true,
	}
	john.Active = false

	fmt.Printf("Name: %s, Age: %d, Active: %t", john.Name, john.Age, john.Active)
}
