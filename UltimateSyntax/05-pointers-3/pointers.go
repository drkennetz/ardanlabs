package main

import "fmt"

// we can also pass the address by value
// this way, we can mutate state outside of our box
// we are sharing across program boundaries

func main() {
	count := 10
	fmt.Println("count:\tValue of count [", count, "]\tAddr Of count [", &count, "]")
	// pass the address of count
	increment(&count)
	fmt.Println("count:\tValue of count [", count, "]\tAddr Of count [", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int
// we pass a copy here, making it value semantic
//go:noinline
func increment(inc *int) {
	//increment value of inc.
	*inc++
	fmt.Println("count:\tValue of inc [", inc, "]\tAddr Of inc [", &inc, "]")
}
