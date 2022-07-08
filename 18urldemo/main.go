package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course_name=reactjs&paymentid=f4c2a12"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qparams)
	fmt.Println(qparams["course_name"])
	fmt.Println(qparams["does_not_exist"])

	for _, val := range qparams {
		fmt.Println("Params:", val)
	}

	partsOfUrl := &url.URL{ // remember to pass by reference!
		Scheme:  "https",
		Host:    "lco.dev:3000",
		Path:    "/learn",
		RawPath: "course=react",
	}

	url2 := partsOfUrl.String()
	fmt.Println(url2)
}
