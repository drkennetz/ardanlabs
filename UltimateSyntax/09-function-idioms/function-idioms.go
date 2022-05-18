package main

import (
	"encoding/json"
	"fmt"
)

// user is a struct type that declares user info
type user struct {
	ID   int
	Name string
}

func main() {
	// retrieve the user profile.
	u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Display the user profile.
	fmt.Printf("%+v\n", *u)
}

// retrieveUser retrieves the user document for the specified user and returns a pointer to a user type value.
func retrieveUser(name string) (*user, error) {
	r, err := getUser(name)

	// if statements should only be used for the negative path
	// else clause should disappear as much as possible
	// column is the happy path. If paths should only be used for the negative path
	// we want to return out of the function as soon as we can for the negative path
	if err != nil {
		// nil is the zero value literal
		return nil, err
	}

	// go has a naked switch which lets you do boolean clauses as a case statement
	//switch {
	//case 1 == 2:
	//}

	// Unmarshal the json document into a value of
	// the user struct type.
	var u user
	// this is being declared inside the if, and only inside the if
	// any variable declared in the scope of a statement will only be in the scope of the statement
	if err := json.Unmarshal([]byte(r), &u); err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUser simulates a web call that returns a json doc for the specified user.
func getUser(name string) (string, error) {
	response := `{"id":1432, "name":"sally"}`
	return response, nil
}
