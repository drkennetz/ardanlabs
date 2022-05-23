package main

import "fmt"

// type is life
// the amount of memory we will read and write

// basic unit of memory is the byte, which is 8 bits
// take 00001010 - this is an arbitrary value without a type

func main() {

	// declare vars at their 0 val
	// we want to let the architecture determine the most efficient int, so we don't declare it without good reason:
	var a int
	fmt.Printf("%T ", a)

	// as opposed to:
	var b int32
	var c int64
	fmt.Printf("%T %T", b, c)

	// we should use var for a 0 value
	// strings can be implemented in many, many ways
	// strings are a 2-word value
	// first word: pointer
	// second word: integer that reps the number of bytes of the string

	// go has conversion over casting
	// casting gives a lot of bugs over time
	// casting can produce bugs
}
