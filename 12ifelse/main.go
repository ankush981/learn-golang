package main

import "fmt"

func main() {
	age := 22

	if age < 18 {
		fmt.Println("Not an adult :(")
	} else {
		fmt.Println("Welcome!")
	}

	// initialize and check (helpful in web requests)
	if num := 12; num > 10 {
		fmt.Println("Above 10!")
	} else {
		fmt.Println("Below 10")
	}
}
