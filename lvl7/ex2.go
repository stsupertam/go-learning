package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname string
	address string
}

func changeMe(p *person,a string) {
	p.address = a
	// (*p).address = a
}

func main() {
	p1 := person{
		firstname: "Test1",
		lastname: "Test2",
		address: "99/99 99",
	}

	fmt.Println(p1)
	changeMe(&p1, "777")
	fmt.Println(p1)
}

