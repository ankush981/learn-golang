package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"source"`
	Password string   `json:"-"` // hide
	Tags     []string `json:"tags"`
}

func main() {
	courses := []course{
		{"Learn React", 12, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Learn Angular", 15, "FreeCodeCamp", "simpli$", []string{"js", "ts"}},
		{"Learn Vue", 9, "Frontend Masters", "login1234", nil},
	}

	coursesJson, err := json.Marshal(courses)

	/*
		We use JSON tags to overcome some annoyances of default marshalling:
		1) The JSON keys are uppercase
		2) Keys carry the same name as in the struct
		3) Fields like `Password` become visible
	*/

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", coursesJson)

	// Now, let's consume the same JSON and unmarshal
	isValidJson := json.Valid(coursesJson)
	fmt.Printf("Is JSON valid while consuming: %v\n", isValidJson)

	if isValidJson {
		var newCourses course
		json.Unmarshal(coursesJson, &newCourses)
		fmt.Printf("%#v\n", newCourses)

		// we can unmarshal to a free-form map as well
		fmt.Println("Unmarshaling to a map:")
		newJson := []byte(`{"coursename":"new course","price":16,"source":"leaf","tags":["tag-1","tag-2"]}`)
		var mapCourses map[string]interface{}
		newData := json.Unmarshal(newJson, &mapCourses)
		fmt.Printf("%s\n", newData) // to show error, if any
		fmt.Printf("%#v\n", mapCourses)
	}
}
