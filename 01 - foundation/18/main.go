package main

import "fmt"

func main() {
	var myVar interface{} = "John Doe"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("The value of res is %v and the result of ok is %v", res, ok)
	res2 := myVar.(int)
	fmt.Printf("The value of res2 is %v", res2)
}
