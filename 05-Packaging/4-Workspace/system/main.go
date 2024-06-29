package main

import (
	"github.com/google/uuid"
	"github.com/torbenish/pos-golang/05-Packaging/3-Replace/math"
)

func main() {

	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
