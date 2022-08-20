package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// here we wish to return a custom error
		return 0, &circleAreaError{"Invalid circle radius", radius}
	}

	return radius * radius * math.Pi, nil
}

// To represent our error (beyond a trivial text message), we need a custom struct
type circleAreaError struct {
	message string
	radius  float64
}

// And now let's implement the standard Error interface to make our struct an Error
// Notice that this implementation is actually trivial for us as it doesn't help us
// get rich information about the error. We do it only to make our custom struct
// can be treated as an error by the Go runtime (this implentation is the reason
// why the function circleArea() can return the struct circleAreaError even though
// its return type is `error`).
func (e *circleAreaError) Error() string {
	return e.message
}

func main() {
	area, err := circleArea(-10.3)
	if err != nil {
		if err, ok := err.(*circleAreaError); ok {
			fmt.Println(err.message, err.radius) // rich error info FTW!
		}
	}
	fmt.Println("Area: ", area)
}
