package main

import "fmt"

func main() {
	salary := map[string]int{
		"Test A": 1000,
		"Test B": 2000,
		"Test C": 3000,
	}
	delete(salary, "Test A")
	salary["Test D"] = 5000

	for nome, salario := range salary {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _,salario := range salary {
		fmt.Printf("O salário é %d\n", salario)
	}
}
