package main

import (
	"fmt"
	"sort"
)

func main() {
	fruitList := []string{"Apple", "Peach", "Mango"}
	fmt.Printf("Type is: %T\n", fruitList)
	fmt.Println("Length is:", len(fruitList))

	fruitList = append(fruitList, "Banana", "Grape")
	fmt.Println(fruitList)
	fmt.Println("Length is:", len(fruitList))   // 4
	fmt.Println("Capacity is:", cap(fruitList)) // 6

	// slice up the slice
	fruitList2 := fruitList[1:5] // allowed, since capacity is 6
	fmt.Println(fruitList2)

	fruitList3 := fruitList[1:3]
	fmt.Println(fruitList3)

	highScores := make([]int, 4) // integer slice of 4
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 445
	highScores[3] = 123

	fmt.Println(highScores)

	// highScores[4] = 650 // if we do this
	// fmt.Println(highScores) // then error, out of bounds

	// but we can append stuff
	highScores = append(highScores, 1001, 1300)
	fmt.Println(highScores)
	fmt.Println("Capacity is:", cap(highScores))

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// Remove a value from a slice (based on index)
	courses := []string{"JavaScript", "React", "Ruby", "Swift", "Python"}
	indexToRemove := 2
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println(courses)
}
