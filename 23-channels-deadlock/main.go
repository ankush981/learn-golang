package main

func main() {
	ch := make(chan int)
	ch <- 2
}

// Since there are no channels to read from `ch`, and since channels are blocking
// by default, the main() goroutine is now blocked forever. The Go runtime detects
// this deadlock and makes the program crash.
