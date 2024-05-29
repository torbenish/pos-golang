package main

import "fmt"

type Address struct {
	publicSpace string
	Number      int
	City        string
	State       string
}

type Person interface {
	Disable()
}

type Enterprise struct {
	Name string
}

func (e Enterprise) Disable() {
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

func Deactivation(person Person) {
	person.Disable()
}

func main() {
	// john := Client{
	// 	Name:   "John",
	// 	Age:    30,
	// 	Active: true,
	// }
	myEnterprise := Enterprise{}

	Deactivation(myEnterprise)
}
