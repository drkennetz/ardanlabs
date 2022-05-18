package main

import "fmt"

// this is not really code you will see, it is just to emphasize the point that you can redeclare and declare in the same line
// when there is a short variable declaration operator being used
// a short variable declaration operator must only construct one new variable to be used, and at least one new variable

type user struct {
	id   int
	name string
}

func main() {
	// declare the error varibale
	var err1 error

	// the short variable declaration operator will declare u and redeclare err1.
	u, err1 := getuser()
	if err1 != nil {
		return
	}

	fmt.Println(u)

	// the short variable declaration operator will redeclare u and declare err2
	u, err2 := getuser()
	if err2 != nil {
		return
	}
	fmt.Println(u)
}

func getuser() (*user, error) {
	return &user{1432, "Betty"}, nil

}
