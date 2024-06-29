package main

import "github.com/torbenish/pos-golang/05-Packaging/3-Replace/math"
func main() {
	
	m := math.NewMath(1, 2)
	println(m.Add())
}