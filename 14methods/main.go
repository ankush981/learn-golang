package main

import "fmt"

func main() {
	phil := User{"Phil", "phil@gmail.com", true, 30, 16}
	fmt.Println(phil)
	fmt.Printf("%+v\n", phil) // + also shows field names
	fmt.Printf("Name: %v, Email: %v\n", phil.Name, phil.Email)
	phil.GetStatus()
	phil.SetNewEmail()
	fmt.Printf("%+v\n", phil)
}

type User struct {
	Name       string
	Email      string
	Status     bool
	Age        int
	privateAge int
}

// becomes a method of User just by passing in User
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// demo of setter
// remember, we need pointers to modify the struct!
func (u *User) SetNewEmail() {
	u.Email = "test@go.dev"
}
