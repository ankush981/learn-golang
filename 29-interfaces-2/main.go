// Internal representation of interfaces
package main

import "fmt"

type Walker interface {
	walk()
}

type Person struct {
	name string
}

func (p Person) walk() {
	// no body
}

func main() {
	p := Person{name: "Ankush"}
	var iface Walker = p
	fmt.Printf("Interface type: %T\n", iface)  // main.Person
	fmt.Printf("Interface value: %v\n", iface) // {Ankush}

	// Now, let's apply type assertion and extract the underlying value
	// from the interface
	p2 := iface.(Person)
	fmt.Println(p2.name)
}

// The output of the above program remains the same even if the interface isn't
// implemented and the Printf statements are run on `p`. Perhaps this just shows
// that nothing magical is going on; it's just the implmenting type's (struct's)
// info being extracted(?).
