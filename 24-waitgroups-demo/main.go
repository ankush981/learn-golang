package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	for _, n := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("%v\n", n)
		time.Sleep(time.Millisecond * 300)
	}
	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
	for _, l := range "abcdefgh" {
		fmt.Printf("%v\n", string(l))
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}

func printBooleans(wg *sync.WaitGroup) {
	for _, b := range []bool{true, true, false, false, true} {
		fmt.Printf("%v\n", b)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3) // We tell the WaitGroup to wait for three Goroutines to finish
	go printNumbers(&wg)
	go printLetters(&wg)
	go printBooleans(&wg)
	wg.Wait()
}

// In the example above, we could've also done `wg.Add(1)` before every `go` call,
// but in the end it doesn't matter. I wrote it this way to show that most tutorials
// out there are just repeating the patterns they come across without deepening one's
// understanding.
