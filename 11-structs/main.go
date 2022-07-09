package main

import "fmt"

func main() {
	phil := User{"Phil", "phil@gmail.com", true, 30}
	fmt.Println(phil)
	fmt.Printf("%+v\n", phil) // + also shows field names
	fmt.Printf("Name: %v, Email: %v\n", phil.Name, phil.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
