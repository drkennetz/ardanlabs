package main

import "fmt"

// value semantics
// sample program to show the bvasic concept of pass by value

// the variable count gives us the "value" of count
// the address "&count" tells us where the box is
func main() {
	count := 10
	fmt.Println("count:\tValue of count [", count, "]\tAddr Of count [", &count, "]")
	increment(count)
	fmt.Println("count:\tValue of count [", count, "]\tAddr Of count [", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int
// we pass a copy here, making it value semantic
//go:noinline
func increment(inc int) {
	//increment value of inc.
	inc++
	fmt.Println("count:\tValue of inc [", inc, "]\tAddr Of inc [", &inc, "]")
}
