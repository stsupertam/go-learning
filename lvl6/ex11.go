package main

import (
	"fmt"
)

func main() {
	a := foo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func foo() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
