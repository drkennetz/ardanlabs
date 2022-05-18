package main

import "fmt"

// because of implicit conversions of kind, you cant have enums in go
// don't try to create aliases for compiler protection, just use the types
// there are trade-offs everywhere
// this is bad:

type color int

// PrintColor should only print if a color is passed
// WRONG - go does kind promotion, so if a color is also an int, the int can be promoted
func PrintColor(c color) {
	fmt.Printf("%T\n", c)
}

func main() {
	var x color = 100
	PrintColor(x)
	// This is not protected by the compiler, 100 gets promoted to color since color is also type int
	PrintColor(100)
}
