package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)

	// Write three values into the channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Println("Sent ", i)
		}
	}()

	fmt.Println("Received ", <-ch)
	fmt.Println("Received ", <-ch)
	fmt.Println("Received ", <-ch)
	fmt.Println("Bye!")
}
