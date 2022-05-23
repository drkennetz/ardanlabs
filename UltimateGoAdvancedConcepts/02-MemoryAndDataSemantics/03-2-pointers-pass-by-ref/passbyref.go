package main

/***
 pointer semantics
 this is a sharing example
 stack looks the same and creates its own input inside that box
 this is a misnomer, this is actually still a pass-by-value, we are just passing a copy of the address which points to the value of count
 the goroutine has shared access to something outside its frame
 we will eventually define guidelines when to use one semantic over the other
***/

func main() {
	count := 10
	println("count:\tValue of [", count, "]\tAddr Of [", &count, "]")

	// Pass the "address of" count - this is still pass by value
	increment(&count)

	// when we pass the address, we have indirect access to the frame above us via the address
	// the goroutine only has direct memory access to the frame it is operating in
	// this gives us efficiency to modify in place, but it has side-effect

	println("count:\tValue of [", count, "]\tAddr Of [", &count, "]")
}

// this is now asking for a pointer to an int
// I want the "address of" the int
func increment(inc *int) {

	// Increment the vale of count that the pointer points to
	*inc++

	println("inc:\tValue of [", inc, "]\tAddr of [", &inc, "]\tValue points to [", *inc, "]")
}
