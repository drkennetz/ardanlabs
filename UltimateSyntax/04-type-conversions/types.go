package main

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141692,
	}

	var ex example
	ex = e
	//
}
