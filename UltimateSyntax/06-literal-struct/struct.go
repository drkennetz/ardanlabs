package main

import "fmt"

// we don't need this anywhere else!
// in this instance I am setting its zero value
var e1 struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e2 := struct {
		flag bool
	}{
		flag: true,
	}
	fmt.Println(e2)
}
