package main

import "fmt"

const (
	// Max integer value on 64 bit architecture
	maxInt = 9223372036854775807

	// much larger value than int64
	bigger = 9223372036854775807243098

	// will NOT compile
	// biggerInt int64 = 9223372036854775807243098
)

func main() {
	fmt.Println(maxInt)
	fmt.Println(bigger)
}
