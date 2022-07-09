package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05")) // this is the syntax instead of YYYY-MM-DD, etc.
	fmt.Println(currentTime.Format("2006-02-01T15:04:05+0:00"))

	// create date/time
	createdDate := time.Date(2020, time.April, 10, 4, 3, 6, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("2006-02-01 Monday"))
}
