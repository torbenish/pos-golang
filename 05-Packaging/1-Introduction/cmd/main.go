package main

import (
	"fmt"
	"github.com/torbenish/pos-golang/05-Packaging/1-Introduction/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	//fmt.Println("Hello, World!")
}
