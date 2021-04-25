package main

import (
	"fmt"
)


func main() {
	a := foo()
	x, y := bar()

	fmt.Println(a)
	fmt.Println(x, y)


}

func foo() int {
	return 999
}

func bar() (int, string) {
	return 999, "Hello"
}