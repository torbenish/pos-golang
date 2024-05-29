package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e)   //tipo
	fmt.Printf("\nO tipo de E é %g", e) //%v=variavel,%g=float
	fmt.Printf("\nO tipo de E é %T", f) //tipo do ID
}
