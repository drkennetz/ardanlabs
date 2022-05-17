package main

import "fmt"

// example represents a type with different fields.
// this is how we declare our own types
// providing compiler amount of memory that needs to be allocated
// provides the representation of the data
// composite type
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// declare a variable of type example set to its zero val
	var e1 example

	fmt.Printf("%+v\n", e1)

	// declare a variable of type example and init using a struct literal
	// can rearrange values
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	// display the field values.
	fmt.Println("Flag: ", e2.flag)
	fmt.Println("Counter: ", e2.counter)
	fmt.Println("Pi: ", e2.pi)

	// empty literal construction -> not empty zero construction
	// it is not the same as zero value

	// we don't want to do partial construction
	// Bill would rather gather all the values in temporary vars before constructing. Example:
	// partial construction
	var e example
	e.flag = true

	// gathering requirements then constructing
	variable := true
	e = example{
		flag: variable,
	}
	fmt.Println(e)
}
