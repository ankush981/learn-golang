package main

import "fmt"

func main() {
	greeter()
	fmt.Println(adder(2, 3))
	fmt.Println(variadicAdder(1, 2, 3))
	fmt.Println(variadicAdder(1, 2, 3, 4, 5))
}

func greeter() {
	fmt.Println("Hello, there!")
}

func adder(n1 int, n2 int) int {
	return n1 + n2
}

func variadicAdder(values ...int) int {
	total := 0
	for _, n := range values {
		total += n
	}
	return total
}
