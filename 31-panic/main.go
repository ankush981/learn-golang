package main

// Because of a its messy, try-catch like semantics, panic (and recover) is best
// avoided in Go. The idiomatic way is to work with Errors. However, for certain
// catastrophic scenarios, it might be a good idea to just panic and exit (failing
// to bind a web server to a port, for example).
// In any case, it doesn't do harm to learn the concept.

func alwaysPanic() {
	panic("Yikes! Something went horribly wrong.") // panic causes the stack trace to be printed
}

func main() {
	alwaysPanic()
}

// Panics cause all defers to execute and return the control to the caller. If you wish, you
// can stop the panic bubbling-up process (much like the catch block in a try-catch)
// but this leads to bad design so we'll skip it.

// One major disadvantage when trying to use recover() is that it works only in the same
// gorutine. So, the launching gouroutine can never recover() from its "child" goroutines,
// making the feature kinda useless.
