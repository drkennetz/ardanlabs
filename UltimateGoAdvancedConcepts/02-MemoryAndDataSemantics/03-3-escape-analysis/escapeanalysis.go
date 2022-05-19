package main

/*** COMPILER ESCAPE ANALYSIS
 go build -gcflags -m=2
 we will use this later for profiling
 the only thing that can tell us why something is allocated is the compiler

 stacks are self-cleaning
 we clean on the way down, and we don't do anything on the way up because we don't know if we will need it again
 something really special about the compiler

 the compiler can perform static code analysis

 ESCAPE ANALYSIS
 the compiler is reading your code at compile time and trying to determine where a value should be constructed in memory
 should it be constructed on the stack, or on the heap?
 where is a value constructed?
 escape analysis is not about construction, it is about how a value is shared
 construction tells you nothing - we don't know where it needs to be constructed
 if we share down, it's all good,
 anytime we share up, we're going to have a problem on the stack
 when we share up, the next function call wipes out what we're sharing
 if the compiler recognizes that a value is going to be shared up, it immediately goes that we cannot construct this on the stack, we must construct it on the heap

 So what does this mean?
 this variable is going to represent a value that is constructed on the heap!
 the variable on the stack, points to a variable on the heap
 we don't have to physically worry where a value is being constructed
 the compilers escape analysis is going to determine where the value should be constructed
 when the heap is involved, this is where the garbage collector comes in
 if performance matters, then these things matter
 a few things will slow down your go program
 internal latencies are things like the garbage collector.
 latencies are what really slow us down
 the compiler will look at how we share something, and determine how we clean

 Guidelines to follow:
 if we are doing construction to a variable, we will be doing value construction
***/

// represents a user in the system
type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passed a copy back to the caller
// this version is going to use value semantics and send a copy back to the caller
// it's being created inside the stack local to the function

/*** value semantic stack
Note: Go-1 and Go-2 and Go-3 are the same routines, they are just annotated to know when the goroutine is where with the flow of our program
STACK
-------------------
M     | Bill-copy | <- Go-3
-------------------
v1    | Bill      | <- Go-1
------------------|
print | Bill-copy | <- Go-2
-------------------
*/

func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V1", &u)

	return u
}

// using pointer semantics, the caller gets shared access to the value that this function constructs

/*** pointer semantic stack
Note: Go-1 and Go-2 and Go-3 are the same routines, they are just annotated to know when the goroutine is where with the flow of our program
STACK
-------------------
M     | Bill-ptr  | <- Go-3
-------------------
v2    | Bill      | <- Go-1
------------------|
print | Bill-ptr  | <- Go-2
-------------------
*/

func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V2", &u)
	return &u
}
