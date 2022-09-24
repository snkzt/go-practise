package main

import "fmt"

// By using recover built-in function, we can recover from panic.
// A recover can stop a panic from aborting the program and
// let it continue with execution instead.
// Example is a server wouldn't want to crash if one of the client
// connections exhibits a critical error.
// Instead, the server would want to close that connection and
// continue serving other clients.: This is what net/http does by
// default for HTTP servers.
func mayPanic() {
	panic("a problem")
}

func main() {
	// Recover must be called within a deferredc function.
	// When the enclosing function of defer painics, defer will activate
	// and a recover call within it will catch the panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd. Error:\n", r)
		}
	}()
	mayPanic()
	// This code won't be run as the execution of the main function
	// stops at mayPanic and resumes at defer.
	fmt.Println("After mayPanic()")
}
