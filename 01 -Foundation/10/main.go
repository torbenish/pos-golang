package main

import "fmt"

func main() {
	
	total := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()
	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
