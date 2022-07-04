package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	welcome := "Welcome to the UserInput program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our pizza: ")
	input, _ := reader.ReadString('\n') // read till end of line
	input = strings.TrimSuffix(input, "\n")
	fmt.Printf("Thanks for rating us for %s stars", input)
	fmt.Println()
}
