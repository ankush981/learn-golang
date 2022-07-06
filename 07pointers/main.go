package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Initial value of pointer is:", ptr)

	myNumber := 13
	ptr = &myNumber
	fmt.Println("After assignment, value of pointer is:", ptr)
	fmt.Println("After assignment, the value referenced to by the pointer is:", *ptr)
}
