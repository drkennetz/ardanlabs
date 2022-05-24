package main

/**
we want 50k goroutines if reasonable and practical
this means each stack will have its own goroutine and we want that stack to be small
starts at 2k
Eventually, you'll make enough function calls and take enough frames off the stack to where the stack must grow
 - go will grow the stack via the compiler - we use contiguous stacks
 - once the stack runs out of memory, go will double the size of the stack 2k -> 4k -> 8k
 - it is a brand new contiguous block of allocation
 - we will need to reframe everything that was originally there
 - when we run out of stack space - these values have to come over
 - values are moving across stacks to another stack
 - the addresses then get adjusted
 &&& imagine the following scenario - 50k goroutines all with their own stacks
   - imagine if we then allowed goroutines to have pointers to stacks
   - everyone would share values with everyone else
   - now suddenly, one of these stacks have to grow!!!
      *** all of these pointers would then need to be adjusted and this would kill performance because of latency
 - One of the constraints of these goroutines is that no goroutines can share values with another goroutine on its stack
 - pointers would be internal to goroutines
 - if two goroutines have to share a value, that value has to be on the heap because we can't share values between stacks
 - value semantics give us that isolation model
 - let's see a live growing stack
*/

// number of element to grow each stack frame.
// run with 10 and then with 1024
const size = 1024

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size of the stack
func stackCopy(s *string, c int, a [size]int) {
	// with a size of 10, we have a 40 byte array and never blow out our stack
	// with an array of size 1024, we have a 40000 byte array
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
	// if we look at indexes 2 and 6, we see the actual address of the string change because it has moved call stacks
}
