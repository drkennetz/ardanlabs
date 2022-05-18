package main

import "fmt"

// functions are first class values in go
// can have named or literal funcs
// when you have multiple defers, the last one defined is the first one that executes

func main() {
	var n int

	// declare an anonymous function and call it.
	// go has closures - n is being declared outside the scope of the func literal, but is being called inside
	// when this function executes, whatever the value of n is at that time, will be printed.
	func() {
		fmt.Println("Direct:", n)
	}()

	// declare an anonymous function and assign it to a variable.
	f := func() {
		fmt.Println("Variable:", n)
	}
	f()

	// defer the call to the anonymous function till after main returns
	defer func() {
		fmt.Println("Defer 1:", n)
	}()

	// set the value of to 3 before the return
	n = 3

	// call the anonymous function through the variable
	f()

	// defer the call to the anonymous function till after main returns.
	defer func() {
		fmt.Println("Defer 2:", n)
	}()

}
