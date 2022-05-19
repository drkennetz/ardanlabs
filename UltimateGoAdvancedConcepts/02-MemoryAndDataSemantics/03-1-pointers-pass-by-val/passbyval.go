package main

/*** make a copy of data as we pass program boundaries

 when your go program starts up
 how many logical processors do you have on your machine?
 physical processors, physical cores
 hardware threads per core

 [P] -> [M] -> [os] -> [hw] processor -> machine -> scheduled by os -> run on hardware
      |
      [G] - goroutine -> application level thread
 every goroutine gets its own contiguous block of memory called a stack which starts at 2k
 os threads also get a stack
 the idea is not only to have high levels of integrity
 it's also about minimizing the amount of resources (like memory) that we use
 50k goroutines makes 50k 2k stacks

 here's our stack, and our goroutine wants to start executing the main function
 every time a goroutine wants to execute a function, we are going across a program boundary
 every time a program crosses over that program boundary, what we're going to do is take a slice of memory off the top of our stack
 this slice is called a frame
 this frame is now a sandbox. The goroutine is now locked into the sandbox.
 The only memory that this goroutine can read and or write is in that frame
 this is protecting the integrity of our program
 every function is essentially a data transformation
 input -> transform -> output
 anytime we use a variable, we are asking for the **value** of that variable -> what's in the box, that integer count in this case
 anytime we use that &, we will say address -> location of the box
 data -> value and the address

 when we call increment, we are crossing a program boundary, and several things are going to happen
 a new frame or box
 the goroutine is going to have to move into this new sandbox
 it can either be passed in, generated, or find its way in through read calls.

 the COST of VALUE SEMANTICS
 inefficiency - there's a lot of inefficiency in value semantics
 I had 4 people I want to give data to
 these people are modifying the data by themselves
 now I want a clear picture of what the data looks like - I have to reassamble it into one view
 it's not efficient to give everyone their own copy
 the opposite of value semantics is pointer semantics
 this allows us to share data across a program boundary
 we gain a lot of efficiency
 there's a big cost
 side-effects - these can cause bugs
***/

func main() {

	// declare variable of type int with val of 10
	count := 10

	println("count:\tValue of [", count, "]\tAddr of [", &count, "]") // bf68

	// Pass the "value of" count across a program boundary
	// make a copy of that value inside the box for count
	// we need to store it - parameters are really a mechanic for getting that value inside that frame
	// make a copy of that integer, and operate on your own copy
	// the goroutine is now operating solely on that copy. It's safe, and on its own.
	// this is value semantics and is the safest way to work with data in our programs
	increment(count) // bf60

	// we can see that the modification only happened in the frame of increment, it does not affect the rest of the program
	println("count:\tValue of [", count, "]\tAddr of [", &count, "]") // bf68

}

//
func increment(inc int) {
	// increment the value of count inside the scope of this new program box
	// in this frame, in this sandbox, it cannot affect any other part of the program
	inc++
	println("inc:\tValue of[", inc, "]\tAddr of [", &inc, "]")
}
