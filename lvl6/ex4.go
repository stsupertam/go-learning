package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname string
	age int
}

func (p person) speak() {
	fmt.Printf("My name is %v.I'm %v year old.\n", p.firstname, p.age)
}
func main() {
	h := person{
		firstname: "Test1",
		lastname: "Test2",
		age: 15,
	}
	h.speak()


}