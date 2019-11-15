package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Define a Person type. Both fields public
	type Person struct {
		Name string
		Age  int
	}

	// Create a Person object and store the pointer to it
	nanodano := &Person{"Bob", 33}
	fmt.Println(nanodano)

	// Structs can also be embedded within other structs.
	// This replaces inheritance by simply storing the
	// data type as another variable
	type Hacker struct {
		Person           Person
		FavoriteLanguage string
	}

	mrrobot := &Hacker{*nanodano, "Go"}
	fmt.Println(mrrobot)
	fmt.Println(mrrobot.Person.Name)
	fmt.Println(mrrobot.FavoriteLanguage)

	// Pointer
	myInt := 334
	myIntPoint := &myInt
	fmt.Println(*myIntPoint)
	fmt.Println(myIntPoint)
	fmt.Println(reflect.TypeOf(myIntPoint))
	fmt.Println(reflect.TypeOf(mrrobot))
	fmt.Println(*mrrobot)
}
