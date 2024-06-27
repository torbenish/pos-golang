package main

import (
	"fmt"
	"github.com/torbenish/pos-golang/05-Packaging/1-Introduction/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}
