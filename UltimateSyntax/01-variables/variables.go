package main

import "fmt"

func main() {
	// declare variables that are set to their zero values
	var a int
	var b string
	var c float64
	var d bool

	// code allocates some memory, reading memory, writing memory
	// to use most efficient code possible, we want to use the most efficient int possible
	// code uses "int" and allows compiler to choose most efficient integer size
	// floats require precision, IEEE 754 spec for float
	// go focuses on integrity to a really high level - even at slight performance loss

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	// declare variables and init.
	// short declaration operator - productivity operator
	// if we init a zero value, we should declare it with var a int
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n", dd, dd)

	// specify type and perform a conversion
	// casting isn't very safe, and many bugs due to casting.

	// instead, it does type conversion
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
