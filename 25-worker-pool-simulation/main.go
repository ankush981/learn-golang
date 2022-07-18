package main

import (
	"fmt"
	"sync"
)

// A simple worker pool implementation using channels and WaitGroups.
// Our workers simply read from a channel of integers from an input
// channel and write their squares to an output channel.

func addJobs(jobsCh chan<- int, wg *sync.WaitGroup) {
	// 100 numbers to crunch (jobs)
	for i := 1; i < 101; i++ {
		jobsCh <- i
	}
	close(jobsCh)
	wg.Done()
}

func worker(jobsCh <-chan int, resultsCh chan<- int, wg2 *sync.WaitGroup) {
	for num := range jobsCh {
		resultsCh <- num * num
	}
	wg2.Done()
}

func addWorkers(jobsCh <-chan int, resultsCh chan<- int, wg *sync.WaitGroup) {
	var wg2 sync.WaitGroup
	// 10 workers
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go worker(jobsCh, resultsCh, &wg2)
	}
	wg2.Wait()
	close(resultsCh)
	wg.Done()
}

func readResults(resultsCh <-chan int, wg *sync.WaitGroup) {
	for sq := range resultsCh {
		fmt.Printf("%v ", sq)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	jobsCh := make(chan int)
	resultsCh := make(chan int)

	wg.Add(1)
	go addJobs(jobsCh, &wg)

	wg.Add(1)
	go addWorkers(jobsCh, resultsCh, &wg)

	wg.Add(1)
	go readResults(resultsCh, &wg)

	wg.Wait()
}
