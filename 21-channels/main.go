package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		// after 2 seconds, write to channel
		ch <- "hello!"
	}()

	fmt.Println("Waiting for a message")
	v := <-ch
	fmt.Println("Received", v)
}
