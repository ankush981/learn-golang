package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var myNumber int = 2
	var myNumberTwo float64 = 4.3
	fmt.Println("The sum is:", myNumber+int(myNumberTwo)) // we lose precision

	// rand.Seed(time.Now().UnixNano()) // random seed
	// fmt.Println(rand.Intn(10))       // 0 to 10

	// generate random number from cryptography
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(myRandomNum)
}
