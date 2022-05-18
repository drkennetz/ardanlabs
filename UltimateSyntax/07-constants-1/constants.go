package main

import "fmt"

func main() {

	// untyped constants
	// constants of a kind can be implicitly converted, constants of a type cannot
	const ui = 12345    // kind: integer
	const uf = 3.131592 // kind: floating-point

	// Typed constants still use the constant type system but their precision is restricted
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	// this will overflow since uint8 has max of 255
	// const myUnit8 uint8 = 1000

	// constant arithmetic supports different kinds.
	// kind promotion is used to dtermine kind in these scenarios

	// variable answer will be of type float64.
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Printf("%T", answer)
	// constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// constant zero will be of kind integer
	const zero = 1 / 3 // KindInt(1) / KindInt(3)
}
