package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println(len(languages)) // 2
	fmt.Println(languages["js"])

	// loop over a map
	for key, value := range languages {
		fmt.Println(key, value)
	}

	// delete stuff from map
	delete(languages, "js")
	fmt.Println(languages)
}
