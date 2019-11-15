package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	for ind, str := range myMap {
		fmt.Printf("%d: %s\n", ind, str)
	}

	switch r := rand.Int(); r {
	case r % 2:
		fmt.Println("Random number r is even.")
	default:
		fmt.Println("Random number r is odd.")
	}
	// r no longer available aftter the switch statement
}
