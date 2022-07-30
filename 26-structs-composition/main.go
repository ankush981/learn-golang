package main

import "fmt"

// Let's model the obligatory blog
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	bio     string // on purpose, we have a field with a clashing name
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author name: ", b.fullName()) // composition in action!
	fmt.Println("Author bio: ", b.bio)
}

func main() {
	author1 := author{
		firstName: "Ankush",
		lastName:  "Thakur",
		bio:       "A nerdy guy",
	}

	blogPost1 := blogPost{
		title:   "Composition in Go!",
		content: "This is how composition can be achieved in Go. We don't use inheritance here!",
		bio:     "Look, a blog with a bio!",
		author:  author1,
	}

	blogPost1.details()
}

// From this, we conclude that Golang has a sort of prototypal inheritance mechanism.
// The `bio` property was found on blogPost and hence pulled from there; otherwise,
// the compiler would've looked "deeper" into the composition hierarchy and
// extracted it from the author.
