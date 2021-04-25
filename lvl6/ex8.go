package main

import (
	"fmt"
)

func main() {
	a := foo()
	a()
}

func foo() func() {
	return func() {
		fmt.Println("Still 0 - 0 T T")
	}
}

