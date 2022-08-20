// An interface is a collection method signatures, but that definition alone
// doesn't do justice to the whole concept. There are many more intricacies to
// cover: the type and value of an interface, the nil interface, its restrctions,
// and so on.

// Types that define the methods in an interface are said to implement that
// interface.

package main

import "fmt"

type Mammal interface {
	eat()
	sleep()
}

type Human struct {
	name string
	age  int
}

func (h Human) eat() {
	fmt.Println(h.name + " is eating")
}

func (h Human) sleep() {
	fmt.Println(h.name + " is sleeping")
}

func main() {
	var m Mammal
	h := Human{
		name: "Bubba",
		age:  50,
	}
	m = h // possible since Human struct implements the Mammal interface
	m.eat()
	h.eat() // same thing
	m.sleep()
	// fmt.Println(m.name) -- not possible; type is defined only in terms of the methods
}

// In the above code, notice how even though we can write `m = h` since both the types
// are equivalent, we can't directly access the Human struct's fields from a Mammal type.
// This shows that type equivalence is tested purely by comparing the methods implemented
// and doesn't extend to the internals of a struct.
