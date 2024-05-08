package main

import "fmt"

type Address struct {
	publicSpace string
	Number      int
	City        string
	State       string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("The client %s was disable", c.Name)
}

func main() {
	john := Client{
		Name:   "John",
		Age:    30,
		Active: true,
	}
	john.Active = false
	john.Disable()
}
