package main

import (
	"fmt"
)

// Person ...
type Person struct {
	Name string
	Age  int
}

// ShowName ...
func (p Person) ShowName() {
	fmt.Println(p.Name)
}

// ChangeName ...
func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func main() {
	nanodano := &Person{"Cornelius", 35}
	nanodano.ShowName()
	nanodano.ChangeName("Marla")
	nanodano.ShowName()
}
