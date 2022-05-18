package main

import "fmt"

// len of an array is part of its type information
func main() {

	//var five [5]int

	// declare an array of 4 integers that is initialized
	//four := [4]int{10, 20, 30, 40}
	//five = four
	// ./array-basics.go: cannot use four (type [4]int) as type [5]int in assignment
	// declare an array of five strings that is initialized to its zero val
	// compiler needs to know the length of an array at compile time

	var fruits [5]string

	// an array is a structure that allocates a contiguous block of memory
	// 80 bytes of contiguous memory because strings are 2 word data structures:
	//  0   1   2   3   4
	//_____________________
	//|nil|nil|nil|nil|nil|
	//---------------------
	//| 0 | 0 | 0 | 0 | 0 |
	//---------------------
	// 16b 16b 16b 16b 16b
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// cost of moving a string around your program only costs moving the pointer because of the two word data structure

	// can't go over the bounds of a range
	// iterate over the array, accessing both the index and the value
	// this is a value semantic - use this when you want to operate on copies
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// iterate over the array, accessing the index and using it to access the value
	for i := range fruits {
		fmt.Println(i, fruits[i])
	}

	// declare an array of 4 integers that is initialized with some values
	numbers := [4]int{10, 20, 30, 40}

	// can do this variadically, although I'm not sure why:
	// numbers := [...]int{10, 20, 30, 40}

	// iterate over the indexes of the array of numbers
	// this is a pointer semantic - use this when you want to operate on the original data
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// go uses "for" for a while loop:
	// for i < len(numbers) {
}
