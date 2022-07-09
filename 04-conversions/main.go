package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our pizza app!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please rate our pizza (1-5): ")
	input, _ := reader.ReadString('\n') // read till end of line
	input = strings.TrimSpace(input)
	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Thanks for rating us", numRating+0.03) // shows Float operations
}
