package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"John": 1000, "Alice": 2000, "Bob": 3000}
	m2 := map[string]float64{"John": 1000.20, "Alice": 2000.3, "Bob": 3000.0}
	m3 := map[string]MyNumber{"John": 1000, "Alice": 2000, "Bob": 3000}
	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(Compare(10, 10))
}
