package main

import (
	"flag"
	"fmt"
)

func main() {
	numProcesses := flag.Int("n", 1, "Number of parallel processes to launch")
	flag.Parse()
	fmt.Printf("Launched %d processes\n", *numProcesses)
}
