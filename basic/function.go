package main

import "fmt"

func variadic(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func checkStatus() (int, error) {
	return 200, nil
}

// Define a type as a function so it can be used
// as a return type
type greeterFunc func(string)

// Generate and return a function
func generateGreetFunc(greeting string) greeterFunc {
	return func(name string) {
		fmt.Printf("%s, %s \n", greeting, name)
	}
}

func main() {
	generateGreetFunc("Hello Boob")

	tajikGreed := generateGreetFunc("Салом")
	tajikGreed("Aleykum")

	statusCode, err := checkStatus()
	fmt.Println(statusCode, err)
}
