package main

import (
	"fmt"
)

// Define a custom type that will be
// used to satisfy the error interface
type customError struct {
	Message string
}

// Satisfy the error interface
// by implementing the Error() function
// wich returns a string
func (c *customError) Error() string {
	return c.Message
}

// Simple function to demonstrate
// how to use custom error
func testFunction() error {
	if true != false {
		return &customError{"Something went wrong."}
	}
	return nil
}

func main() {
	err := testFunction()
	if err != nil {
		fmt.Println(err)
	}
}
