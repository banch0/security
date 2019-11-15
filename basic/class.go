package main

import (
	"fmt"
	"reflect"
)

// Person ...
type Person struct {
	Name string
	Age  int
}

// Doctor ...
type Doctor struct {
	Person         Person
	Specialization string
}

// This is constructor
func creareNewPerson() *Person {
	return &Person{
		Name: "Anonymous",
		Age:  23,
	}
}

func main() {
	// This is inheritance
	nanodano := Person{
		Name: "NanoDano",
		Age:  22,
	}
	// This is inheritance
	drDano := Doctor{
		Person:         nanodano,
		Specialization: "Hacking",
	}
	fmt.Println(reflect.TypeOf(nanodano))
	fmt.Println(nanodano)
	fmt.Println(reflect.TypeOf(drDano))
	fmt.Println(drDano)

	// Polymorphism
	// You can youse with interfaces

	// This is constructor
	p := creareNewPerson()
	fmt.Println(p)
}
