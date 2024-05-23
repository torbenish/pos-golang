package main

import (
	"fmt"
	"pos-golang/matematics"
)

func main() {
	s := matematics.Sum(10, 20)
	car := matematics.Car{Brand: "Fiat"}
	fmt.Println(car.Brand)
	fmt.Println(car.Walk())
	fmt.Println("Result: ", s)
	fmt.Println(matematics.A)
}
