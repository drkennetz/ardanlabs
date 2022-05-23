package main

import "fmt"

// structs define the data you need. you have to model the data yourself
// type provides 2 pieces of information to the compiler:
// 1. How much memory you need to allocate
// 2. what that memory represents

// 8 contiguous bytes to define this struct
// bool | int16 | float32
//   1      2       4
// this is only 7 bytes, but the machine has to do something called alignments
// the alignments are important to many mechanical sympathies inside the machine
//
// understand the cost
// this is an 8 byte value because alignments are going to push an extra byte inside the struct
// this extra byte is called the padding byte, and it will go between the bool and the int16

// bool | padding | int16 | float32
//  1        1        2        4
// from a hardware perspective, we need to make sure memory is always aligning along proper boundaries
// we are creating mechanical sympathies with the hardware
// minimize the operations to bring data in and out of the hardware
// we want to avoid 2 byte values falling across word boundaries, so we say
// all 2 byte values fall within 2 byte alignments
// we can guarantee that those 2 bytes fall within that machine word evenly
// 4 byte value requires a 4 byte alignment

// the struct itself has to fall on an alignment that is the size of the largest field

// some mental models for padding
// type example struct {
//   flag bool - 1
//   [7]byte padding -> so the int64 can fall on an 8byte alignment
//   counter int64 - 8
//   pi float32 - 4
//   [4]byte padding -> so the struct can fall on the 8 byte alignment
// }

// type example struct {
//   flag bool - 1
//   [3]byte padding -> so the int32 can fall on a 4byte alignment
//   counter int32 - 4
//   pi float32 - 4
//   [0]byte padding -> this already falls on 4 byte alignment
// }

// we do not do this type of microalignment unless there is a constraint, or a profile that our profiler is showing us
// one simple model is to align your fields from largest to smallest

// example represents a composite type with different fields
type example struct {
	pi      float32
	counter int16
	flag    bool
}

func main() {
	// declare a var of type example set to its zero value
	var e1 example

	fmt.Printf("%+v\n", e1)
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println("Flag: ", e2.flag)
	fmt.Println("Counter: ", e2.counter)
	fmt.Println("Pi: ", e2.pi)
}
